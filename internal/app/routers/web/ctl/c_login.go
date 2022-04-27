package ctl

import (
	"github.com/LyricTian/captcha"
	"github.com/gin-gonic/gin"
	"go-gin/internal/app/bll"
	"go-gin/internal/app/config"
	"go-gin/internal/app/errors"
	"go-gin/internal/app/ginplus"
	"go-gin/internal/app/schema"

	"go-gin/pkg/auth"
	"go-gin/pkg/logger"
	"time"
)

// NewLogin 创建login控制器
func NewLogin(BLogin bll.ILogin) *Login {
	return &Login{
		LoginBll: BLogin,
	}
}

// Login Login
// @Name 登录
// @Description 登录接口
type Login struct {
	LoginBll bll.ILogin
	Auth     auth.Auther
}

// GetCaptcha 获取验证码信息
// @Summary 获取验证码信息
// @Success 200 schema.LoginCaptcha
// @Router GET /web/v1/pub/login/captchaid
func (a *Login) GetCaptcha(c *gin.Context) {
	item, err := a.LoginBll.GetCaptcha(ginplus.NewContext(c), config.GetGlobalConfig().Captcha.Length)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// ResCaptcha 响应图形验证码
// @Summary 响应图形验证码
// @Param id query string true "验证码ID"
// @Param reload query string false "重新加载"
// @Success 200 file "图形验证码"
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /web/v1/pub/login/captcha
func (a *Login) ResCaptcha(c *gin.Context) {
	captchaID := c.Query("id")
	if captchaID == "" {
		ginplus.ResError(c, errors.ErrInvalidRequestParameter)
		return
	}

	if c.Query("reload") != "" {
		if !captcha.Reload(captchaID) {
			ginplus.ResError(c, errors.ErrInvalidRequestParameter)
			return
		}
	}

	cfg := config.GetGlobalConfig().Captcha
	err := a.LoginBll.ResCaptcha(ginplus.NewContext(c), c.Writer, captchaID, cfg.Width, cfg.Height)
	if err != nil {
		ginplus.ResError(c, err)
	}
}

// Login 用户登录
// @Summary 用户登录
// @Param body body schema.LoginInfo false
// @Success 200 schema.LoginTokenInfo
// @Failure 400 schema.HTTPError "{error:{code:400,message:无效的请求参数}}"
// @Failure 500 schema.HTTPError "{error:{code:500,message:服务器错误}}"
// @Router POST /web/v1/pub/login
func (a *Login) Login(c *gin.Context) {
	var item schema.LoginInfo
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
	}
	if item.Type != "" {
		if item.Type == "password" {
			a.HandlePassWord(c, item)
		}
	}

}

func (a *Login) HandlePassWord(c *gin.Context, item schema.LoginInfo) {

	ctx := ginplus.NewContext(c)

	if !captcha.VerifyString(item.CaptchaID, item.CaptchaCode) {
		ginplus.ResError(c, errors.ErrLoginInvalidVerifyCode)
		return
	}

	user, err := a.LoginBll.Verity(ctx, item.UserName, item.PassWord)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	userID := user.RecordID
	// 将用户ID放入上下文
	ginplus.SetUserID(c, userID)

	tokenInfo, err := a.LoginBll.GenerateToken(ctx, userID)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	// 登录用户登录时间
	user.LastLoginTime = time.Now()
	//_ = a.userBll.Update(ctx, userID, *user)
	//// 插入更新用户的系统消息
	//_ = a.userNotify.InsertUserNotify(ctx, user)
	data := &schema.LoginUserInfo{
		LoginTokenInfo: *tokenInfo,
		UserInfo:       *user,
	}
	logger.StartSpan(ginplus.NewContext(c), logger.SetSpanTitle("用户登录"), logger.SetSpanFuncName("Login")).Infof("登入系统")
	ginplus.ResSuccess(c, data)
}

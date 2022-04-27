package internal

import (
	"context"
	"github.com/LyricTian/captcha"
	"go-gin/internal/app/config"
	"go-gin/internal/app/errors"
	"go-gin/internal/app/model"
	"go-gin/internal/app/schema"
	"go-gin/pkg/auth"
	"go-gin/until"
	"net/http"
)

// NewLogin 创建login
func NewLogin(mLogin model.IUser, a auth.Auther) *Login {
	return &Login{
		LoginModel: mLogin,
		Auth:       a,
	}
}

// Login 登录实例
type Login struct {
	LoginModel model.IUser
	Auth       auth.Auther
}

// GetRootUser 获取root用户
func GetRootUser() *schema.UserInfo {
	user := config.GetGlobalConfig().Root
	return &schema.UserInfo{
		RecordID: user.UserName,
		UserName: user.UserName,
		PassWord: until.MD5HashString(user.Password),
	}
}

// GetCaptcha 获取图形验证码信息
func (a *Login) GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error) {
	captchaID := captcha.NewLen(length)
	item := &schema.LoginCaptcha{
		CaptchaID: captchaID,
	}
	return item, nil
}

// ResCaptcha 生成并响应图形验证码
func (a *Login) ResCaptcha(ctx context.Context, w http.ResponseWriter, captchaID string, width, height int) error {
	err := captcha.WriteImage(w, captchaID, width, height)
	if err != nil {
		if err == captcha.ErrNotFound {
			return errors.ErrNotFound
		}
		return errors.WithStack(err)
	}
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	return nil
}

// Verity 登录验证
func (a *Login) Verity(ctx context.Context, userName, password string) (*schema.UserInfo, error) {
	result, err := a.LoginModel.Query(ctx, schema.UserQueryParams{
		UserName: userName,
	})
	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, errors.ErrInvalidUserName
	}

	item := result.Data[0]

	// 检查是否是超级用户
	root := GetRootUser()
	if userName == root.UserName {
		if root.PassWord == password {
			return item, nil
		}
		return nil, errors.ErrInvalidUserName
	}

	if item.PassWord != until.SHA1HashString(password) {
		return nil, errors.ErrInvalidPassword
	} else if item.Status != 1 {
		return nil, errors.ErrUserDisable
	}

	return item, nil
}

// GenerateToken 生成令牌
func (a *Login) GenerateToken(ctx context.Context, userID string) (*schema.LoginTokenInfo, error) {
	tokenInfo, err := a.Auth.GenerateToken(userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	item := &schema.LoginTokenInfo{
		AccessToken: tokenInfo.GetAccessToken(),
		TokenType:   tokenInfo.GetTokenType(),
		ExpiresAt:   tokenInfo.GetExpiresAt(),
	}
	return item, nil
}

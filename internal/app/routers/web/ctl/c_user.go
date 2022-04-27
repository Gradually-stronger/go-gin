package ctl

import (
	"github.com/gin-gonic/gin"
	"go-gin/internal/app/bll"
	"go-gin/internal/app/ginplus"
	"go-gin/internal/app/schema"
)

// NewUser 创建login控制器
func NewUser(BUser bll.IUser) *User {
	return &User{
		UserBll: BUser,
	}
}

// User User
// @Name 用户
// @Description 用户接口
type User struct {
	UserBll bll.IUser
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.UserInfo true
// @Success 200 schema.UserInfo
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /web/v1/users
func (a *User) Create(c *gin.Context) {
	var item schema.UserInfo
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	item.Creator = ginplus.GetUserID(c)
	nitem, err := a.UserBll.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem.CleanSecure())
}

// Get 查询指定数据
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 schema.UserInfo
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /web/v1/users
func (a *User) Get(c *gin.Context) {

	UserId := ginplus.GetUserID(c)
	item, err := a.UserBll.Get(ginplus.NewContext(c), UserId, true)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item.CleanSecure())
}

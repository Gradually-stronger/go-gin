package bll

import (
	"context"
	"go-gin/internal/app/schema"
	"net/http"
)

type ILogin interface {
	//GetCaptcha 获取图形验证码信息
	GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error)

	//ResCaptcha 生成并响应图形验证码
	ResCaptcha(ctx context.Context, w http.ResponseWriter, captchaID string, width, height int) error

	//Verity 登录验证
	Verity(ctx context.Context, userName, password string) (*schema.UserInfo, error)

	//GenerateToken 生成令牌
	GenerateToken(ctx context.Context, userID string) (*schema.LoginTokenInfo, error)
}

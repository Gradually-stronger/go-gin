package schema

//LoginInfo 密码验证码登录所需要参数
type LoginInfo struct {
	UserName    string `json:"user_name" swaggo:"false,用户名"`
	PassWord    string `json:"password"  swaggo:"false,密码"`
	CaptchaID   string `json:"captcha_id"  swaggo:"false,验证码ID"`
	CaptchaCode string `json:"captcha_code" swaggo:"false,验证码"`
	Type        string `json:"type" binding:"required" swaggo:"true,登录状态"`
}

//WechatLogin 微信登录
type WechatLogin struct {
	Code       string `json:"code" swaggo:"true,微信code"`
	OpenedId   string `json:"opened_id" swaggo:"true,微信opened_id"`
	NickName   string `json:"nick_name" swaggo:"true,微信昵称"`
	Sex        int    `json:"sex" swaggo:"true,性别"`
	Province   string `json:"province" swaggo:"true,省份"`
	City       string `json:"city" swaggo:"true, 城市"`
	UnionId    string `json:"union_id" swaggo:"true, 微信唯一id"`
	Country    string `json:"country" swaggo:"true, 国家"`
	HeadImgUrl string `json:"head_img_url" swaggo:"true,微信头像"`
}

// LoginCaptcha 登录验证码
type LoginCaptcha struct {
	CaptchaID string `json:"captcha_id" swaggo:"false,验证码ID"`
}

// LoginTokenInfo 登录令牌信息
type LoginTokenInfo struct {
	AccessToken string `json:"access_token" swaggo:"true,访问令牌"`
	TokenType   string `json:"token_type" swaggo:"true,令牌类型"`
	ExpiresAt   int64  `json:"expires_at" swaggo:"true,令牌到期时间"`
}

type LoginUserInfo struct {
	LoginTokenInfo
	UserInfo
}

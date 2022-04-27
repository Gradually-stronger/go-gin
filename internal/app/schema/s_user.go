package schema

import "time"

//UserInfo 用户信息
type UserInfo struct {
	RecordID      string    `json:"record_id" swaggo:"false,记录id"`
	UserName      string    `json:"user_name" swaggo:"true,名称"`
	PassWord      string    `json:"password" swaggo:"true,密码"`
	Age           string    `json:"age" swaggo:"false,年龄"`
	Mobile        string    `json:"mobile" swaggo:"true,手机号"`
	Photo         string    `json:"photo" swaggo:"false,头像"`
	OpenedId      string    `json:"opened_id" swaggo:"false,微信opened_id"`
	NickName      string    `json:"nick_name" swaggo:"false,微信昵称"`
	Sex           int       `json:"sex" swaggo:"false,性别"`
	Province      string    `json:"province" swaggo:"false,省份"`
	City          string    `json:"city" swaggo:"false, 城市"`
	UnionId       string    `json:"union_id" swaggo:"false, 微信唯一id"`
	Country       string    `json:"country" swaggo:"false, 国家"`
	HeadImgUrl    string    `json:"head_img_url" swaggo:"false,微信头像"`
	Status        int       `json:"status" swaggo:"false,用户状态 1.启用 2.禁用"`
	IdCard        string    `json:"id_card" swaggo:"false, 身份认证正面"`
	IdCardNo      string    `json:"id_card_no" swaggo:"false, 身份认证反面"`
	RealStatus    string    `json:"real_status" swaggo:"false, 用户认证状态 auth:认证"`
	Email         string    `json:"email" swaggo:"false, 用户邮箱"`
	LastLoginTime time.Time `json:"last_login_time" swaggo:"false,最后登录时间"`
	Creator       string    `json:"creator" swaggo:"false, 创建人"`
}

type UserQueryParams struct {
	UserName     string   // 用户名
	LikeUserName string   // 用户名(模糊查询)
	Status       int      // 用户状态(1:启用 2:停用)
	RecordIDs    []string // 记录ID列表
	UserID       string   //用户ID 限制查询权限
}

// Result 查询结果
type Result struct {
	Data       Users
	PageResult *PaginationResult
}

// Users 用户对象列表
type Users []*UserInfo

// UserQueryOptions 查询可选参数项
type UserQueryOptions struct {
	PageParam *PaginationParam // 分页参数
}

// CleanSecure 清理安全数据
func (a *UserInfo) CleanSecure() *UserInfo {
	a.PassWord = "******"
	return a
}

package entity

import (
	"context"
	"go-gin/internal/app/schema"
	"go-gin/pkg/gormplus"
	"time"
)

// GetUserDB 获取用户存储
func GetUserDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return getDBWithModelPlus(ctx, defDB, User{})
}

// SchemaUser 用户对象
type SchemaUser schema.UserInfo

// ToUser 转换为用户实体
func (a SchemaUser) ToUser() *User {
	item := &User{
		RecordID:      a.RecordID,
		UserName:      &a.UserName,
		PassWord:      &a.PassWord,
		Age:           &a.Age,
		Mobile:        &a.Mobile,
		Photo:         &a.Photo,
		OpenedId:      &a.OpenedId,
		NickName:      &a.NickName,
		Sex:           &a.Sex,
		Province:      &a.Province,
		City:          &a.City,
		UnionId:       &a.UnionId,
		Country:       &a.Country,
		HeadImgUrl:    &a.HeadImgUrl,
		Status:        &a.Status,
		IdCard:        &a.IdCard,
		IdCardNo:      &a.IdCardNo,
		RealStatus:    &a.RealStatus,
		Email:         &a.Email,
		LastLoginTime: a.LastLoginTime,
	}
	return item
}

// User 用户实体
type User struct {
	Model
	RecordID      string    `gorm:"column:record_id;size:36;index;comment:'记录ID'"`
	UserName      *string   `gorm:"column:user_name;size:50;comment:'true,名称'"`
	PassWord      *string   `gorm:"column:password;size:100;comment:'false,密码'"`
	Age           *string   `gorm:"column:age;size:50;comment:'false,年龄'"`
	Mobile        *string   `gorm:"column:mobile;size:25;comment:'true,手机号'"`
	Photo         *string   `gorm:"column:photo;size:150;comment:'false,头像'"`
	OpenedId      *string   `gorm:"column:opened_id;size:50;comment:'false,微信opened_id'" `
	NickName      *string   `gorm:"column:nick_name;size:255;comment:'false,微信昵称'"`
	Sex           *int      `gorm:"column:sex;size:0;comment:'false,性别'"`
	Province      *string   `gorm:"column:province;size:255;comment:'false,省份'"`
	City          *string   `gorm:"column:city;size:255; comment:'false, 城市'"`
	UnionId       *string   `gorm:"column:union_id;size:100;comment:'false, 微信唯一id'"`
	Country       *string   `gorm:"column:country;size:255;comment:'false, 国家'"`
	HeadImgUrl    *string   `gorm:"column:head_img_url;size:100;comment:'false,微信头像'"`
	Status        *int      `gorm:"column:status;size:50;comment:'false,用户状态 1.启用 2.禁用'"`
	IdCard        *string   `gorm:"column:id_card;size:50;comment:'false, 身份认证正面'"`
	IdCardNo      *string   `gorm:"column:id_card_no;size:50;comment:'false, 身份认证反面'"`
	RealStatus    *string   `gorm:"column:real_status;size:20;comment:'false, 用户认证状态 auth:认证'"`
	Email         *string   `gorm:"column:email;size:50;comment:'false, 用户邮箱'"`
	LastLoginTime time.Time `gorm:"column:last_login_time;comment:'false, 最后登录时间'"`
}

func (a User) String() string {
	return toString(a)
}

// TableName 表名
func (a User) TableName() string {
	return a.Model.TableName("user")
}

// ToSchemaUser 转换为用户对象
func (a User) ToSchemaUser() *schema.UserInfo {
	item := &schema.UserInfo{
		RecordID:      a.RecordID,
		UserName:      *a.UserName,
		PassWord:      *a.PassWord,
		Age:           *a.Age,
		Mobile:        *a.Mobile,
		Photo:         *a.Photo,
		OpenedId:      *a.OpenedId,
		NickName:      *a.NickName,
		Sex:           *a.Sex,
		Province:      *a.Province,
		City:          *a.City,
		UnionId:       *a.UnionId,
		Country:       *a.Country,
		HeadImgUrl:    *a.HeadImgUrl,
		Status:        *a.Status,
		IdCard:        *a.IdCard,
		IdCardNo:      *a.IdCardNo,
		RealStatus:    *a.RealStatus,
		Email:         *a.Email,
		LastLoginTime: a.LastLoginTime,
	}
	return item
}

// Users 用户实体列表
type Users []*User

// ToSchemaUsers 转换为用户对象列表
func (a Users) ToSchemaUsers() []*schema.UserInfo {
	list := make([]*schema.UserInfo, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUser()
	}
	return list
}

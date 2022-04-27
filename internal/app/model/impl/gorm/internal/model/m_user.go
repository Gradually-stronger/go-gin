package model

import (
	"context"
	"go-gin/internal/app/errors"
	"go-gin/internal/app/model/impl/gorm/internal/entity"
	"go-gin/internal/app/schema"
	"go-gin/pkg/gormplus"
)

// NewUser 创建NewUser存储实例
func NewUser(db *gormplus.DB) *User {
	return &User{db}
}

// User 存储db
type User struct {
	db *gormplus.DB
}

func (a *User) getQueryOption(opts ...schema.UserQueryOptions) schema.UserQueryOptions {
	var opt schema.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *User) Query(ctx context.Context, params schema.UserQueryParams, opts ...schema.UserQueryOptions) (*schema.Result, error) {
	db := entity.GetUserDB(ctx, a.db).DB

	if v := params.UserName; v != "" {
		db = db.Where("user_name=?", v)
	}
	if v := params.LikeUserName; v != "" {
		db = db.Where("user_name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.Users
	pr, err := WrapPageQueryNC(db, opt.PageParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.Result{
		PageResult: pr,
		Data:       list.ToSchemaUsers(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *User) Get(ctx context.Context, recordID string) (*schema.UserInfo, error) {
	var item entity.User
	ok, err := a.db.FindOne(entity.GetUserDB(ctx, a.db).Where("record_id=?", recordID), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaUser(), nil
}

// Create 创建数据
func (a *User) Create(ctx context.Context, item schema.UserInfo) error {

	sitem := entity.SchemaUser(item)
	result := entity.GetUserDB(ctx, a.db).Create(sitem.ToUser())
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

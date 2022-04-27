package model

import (
	"context"
	"go-gin/internal/app/schema"
)

type IUser interface {
	// 查询数据
	Query(ctx context.Context, params schema.UserQueryParams, opts ...schema.UserQueryOptions) (*schema.Result, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string) (*schema.UserInfo, error)
	// 创建数据
	Create(ctx context.Context, item schema.UserInfo) error
}

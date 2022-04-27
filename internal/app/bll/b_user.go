package bll

import (
	"context"
	"go-gin/internal/app/schema"
)

type IUser interface {
	//Create 创建数据
	Create(ctx context.Context, item schema.UserInfo) (*schema.UserInfo, error)

	// 查询指定数据
	Get(ctx context.Context, recordID string, includePositon bool) (*schema.UserInfo, error)
}

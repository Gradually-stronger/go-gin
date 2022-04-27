package internal

import (
	"context"
	icontext "go-gin/internal/app/context"
	"go-gin/internal/app/model"
)

// GetUserID 获取用户ID
func GetUserID(ctx context.Context) string {
	userID, _ := icontext.FromUserID(ctx)
	return userID
}

// TransFunc 定义事务执行函数
type TransFunc func(context.Context) error

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, transModel model.ITrans, fn TransFunc) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}
	trans, err := transModel.Begin(ctx)
	if err != nil {
		return err
	}

	err = fn(icontext.NewTrans(ctx, trans))
	if err != nil {
		_ = transModel.Rollback(ctx, trans)
		return err
	}
	return transModel.Commit(ctx, trans)
}

// CheckIsRootUser 检查是否是root用户
func CheckIsRootUser(ctx context.Context, userIDs ...string) bool {
	if len(userIDs) > 0 {
		return GetRootUser().RecordID == userIDs[0]
	}
	return GetRootUser().RecordID == GetUserID(ctx)
}

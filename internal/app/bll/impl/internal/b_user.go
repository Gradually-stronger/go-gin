package internal

import (
	"context"
	"go-gin/internal/app/errors"
	"go-gin/internal/app/model"
	"go-gin/internal/app/schema"
	"go-gin/until"
	"time"
)

// NewUser 创建user
func NewUser(
	mUser model.IUser,
	trans model.ITrans,
) *User {
	return &User{
		UserModel:  mUser,
		TransModel: trans,
	}
}

// User 用户实例
type User struct {
	UserModel  model.IUser
	TransModel model.ITrans
}

func (a *User) checkUserName(ctx context.Context, userName string) error {
	if userName == GetRootUser().UserName {
		return errors.ErrResourceExists
	}

	result, err := a.UserModel.Query(ctx, schema.UserQueryParams{
		UserName: userName,
	}, schema.UserQueryOptions{
		PageParam: &schema.PaginationParam{PageSize: -1},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.ErrResourceExists
	}
	return nil
}

// Create 创建用户数据
func (a *User) Create(ctx context.Context, req schema.UserInfo) (*schema.UserInfo, error) {
	err := a.checkUserName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	if req.PassWord == "" {
		return nil, errors.ErrUserNotEmptyPwd
	}
	var item schema.UserInfo

	if err := ExecTrans(ctx, a.TransModel, func(c context.Context) error {
		//	//创建用户
		item.RecordID = until.MustUUID()
		item.PassWord = until.SHA1HashString(req.PassWord)
		item.UserName = req.UserName
		item.Mobile = req.Mobile
		item.Email = req.Email
		item.Status = 1
		item.LastLoginTime = time.Now()
		//
		_ = a.UserModel.Create(ctx, item)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return a.getUpdate(ctx, item.RecordID)
}

func (a *User) getUpdate(ctx context.Context, recordID string) (*schema.UserInfo, error) {
	nitem, err := a.Get(ctx, recordID, true)
	if err != nil {
		return nil, err
	}

	//err = a.LoadPolicy(ctx, *nitem)
	//if err != nil {
	//	return nil, err
	//}
	return nitem, nil
}

// Get 查询指定数据
func (a *User) Get(ctx context.Context, recordID string, includeRole bool) (*schema.UserInfo, error) {
	var item *schema.UserInfo
	var err error

	item, err = a.UserModel.Get(ctx, recordID)

	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	//屏蔽密码
	item.CleanSecure()

	return item, nil
}

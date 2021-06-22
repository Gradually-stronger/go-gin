package entity

import (
	"context"
	"fmt"
	"time"

	icontext "go-gin/internal/app/context"
	"go-gin/pkg/gormplus"
	"go-gin/until"

	"github.com/jinzhu/gorm"
)

// 表名前缀
var tablePrefix string

// SetTablePrefix 设定表名前缀
func SetTablePrefix(prefix string) {
	tablePrefix = prefix
}

// GetTablePrefix 获取表名前缀
func GetTablePrefix() string {
	return tablePrefix
}

// Model base model
type Model struct {
	ID        uint       `gorm:"column:id;primary_key;auto_increment;"`
	CreatedAt time.Time  `gorm:"column:created_at;comment:'创建时间'"`
	UpdatedAt time.Time  `gorm:"column:updated_at;comment:'更新时间'"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;comment:'删除时间'"` // 此处如果不是*time.Time指针类型，数据库内会默认插入时间的0值，而不是NULL
	OpUser    string     `gorm:"column:op_user;size:36;comment:'操作人'"`
	OpDesc    string     `gorm:"column:op_desc;size:1024;comment:'操作原因'"`
}

type MarketingModel struct {
	ID        uint       `gorm:"column:id;primary_key;auto_increment;unsigned"`
	CreatedAt time.Time  `gorm:"column:created_at;comment:'创建时间'"`
	UpdatedAt time.Time  `gorm:"column:updated_at;comment:'更新时间'"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;comment:'删除时间'"` // 此处如果不是*time.Time指针类型，数据库内会默认插入时间的0值，而不是NULL
	OpUser    string     `gorm:"column:op_user;size:36;comment:'操作人'"`
}

type OperationInfo struct {
}

// TableName table name
func (Model) TableName(name string) string {
	return fmt.Sprintf("%s%s", "p_", name)
}

func toString(v interface{}) string {
	return until.JSONMarshalToString(v)
}

func getDBPlus(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	trans, ok := icontext.FromTrans(ctx)
	if ok {
		db, ok := trans.(*gormplus.DB)
		if ok {
			return db
		}
	}
	return defDB
}

func getDBWithModelPlus(ctx context.Context, defDB *gormplus.DB, m interface{}) *gormplus.DB {
	return gormplus.Wrap(getDBPlus(ctx, defDB).Model(m))
}

func getDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := icontext.FromTrans(ctx)
	if ok {
		db, ok := trans.(*gormplus.DB)
		if ok {
			return db.GetDB()
		}
	}
	return defDB
}

func getDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return getDB(ctx, defDB).Model(m)
}

// GetDB 获取数据库
func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDB(ctx, defDB)
}

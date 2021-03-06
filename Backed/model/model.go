package model

import (
	"ITBat/pkg/logging"
	"ITBat/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID int64 `gorm:"column:Id;primary_key" json:"id"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func init() {
	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DatabaseCfg.USER, setting.DatabaseCfg.PASSWORD, setting.DatabaseCfg.HOST, setting.DatabaseCfg.NAME)
	DB, err = gorm.Open(setting.DatabaseCfg.TYPE, conn)
	if err != nil {
		logging.DebugLog(err)
	}

	// 全局禁用表名复数
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer DB.Close()
}

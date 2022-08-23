package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"

	gormless "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

// Connect 数据库
func Connect(dbConfig gorm.Dialector, _logger gormless.Interface) {

	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

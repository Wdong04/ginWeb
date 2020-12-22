package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB 数据库操作对象
	DB *gorm.DB
)

// InitMySQL 连接MySQL
func InitMySQL() (err error) {
	dsn := "root:mysql123@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
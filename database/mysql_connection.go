package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jira/model"
	"log"
)

var DB *gorm.DB

func Init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/jira?charset=utf8mb4&parseTime=True&loc=Local"
	// 使用 gorm.Open 替代 sql.Open
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接池失败: %v", err)
	}
	fmt.Println("MySQL資料庫連接成功")
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	DB = db

	// 自動根據struct創建表數據
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Println("初始化資料表有問題...")
		return
	}
}

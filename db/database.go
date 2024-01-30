package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"workspace/logs/modules/monitoringErrors"
	"workspace/logs/modules/settings"
)

var (
	DB *gorm.DB
)

// DbInit 初始化数据库
func DbInit() {

	// 设置数据库连接信息
	dsn := "root:Qsj.0228@tcp(localhost:3306)/monitoring?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模型（创建表）
	err = db.AutoMigrate(&monitoringErrors.FrontendErrorReport{}, &settings.LogSettings{})
	if err != nil {
		panic("failed to migrate database")
	}

	fmt.Println("Database connection successful")

	DB = db
}

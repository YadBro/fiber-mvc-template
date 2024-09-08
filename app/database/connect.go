package database

import (
	"fmt"

	"github.com/YadBro/fiber-mvc-template/app/models"
	"github.com/YadBro/fiber-mvc-template/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	dsn := env.GetEnv("DB_USERNAME", "root") + ":" + env.GetEnv("DB_PASSWORD", "") + "@tcp(127.0.0.1:3306)/" + env.GetEnv("DB_NAME", "test_db") + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.Menu{})
	fmt.Println("Database Migrated", DB)
}

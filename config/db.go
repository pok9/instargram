package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DATABASE_CONNECTION")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// db.Logger.LogMode(1)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

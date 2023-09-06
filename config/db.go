package config

import (
	"fmt"
	"os"

	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_CONNECTION"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Photo{})

	if err != nil {
		panic(err)
	}

	DB = db
}

package database

import (
	"fmt"
	"uts/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "username:password@tcp(127.0.0.1:3306)/database_name?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
	fmt.Println("Failed to connect to database:", err)
	return
	}

	// Set global database connection
	DB = db
	DB.AutoMigrate(models.User{} )
}

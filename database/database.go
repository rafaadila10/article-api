package database

import (
	"article-api/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/article?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connect to DB:", err)
	}

	// Auto migrate table posts
	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("failed migrate:", err)
	}

	fmt.Println("DB connected & migrated successfully!")
}

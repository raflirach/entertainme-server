package config

import (
	"entertaime-server/models/movie"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/entertainme?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed!")
	}

	db.AutoMigrate(&movie.Movie{})
	database = db
}

func GetDatabase() (DB *gorm.DB) {
	return database
}

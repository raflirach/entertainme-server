package config

import (
	"entertaime-server/models/movie"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() {
	dsn := "admin:123123123@tcp(first-database.cobquwjdlho5.ap-southeast-1.rds.amazonaws.com)/entertainme?charset=utf8mb4&parseTime=True&loc=Local"
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

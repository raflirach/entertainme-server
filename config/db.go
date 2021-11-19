package config

import (
	"entertaime-server/models/movie"
	"entertaime-server/models/series"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() {
	godotenv.Load()
	db_username := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/entertainme?charset=utf8mb4&parseTime=True&loc=Local",
		db_username, db_pass, db_host,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed!")
	}

	db.AutoMigrate(&movie.Movie{})
	db.AutoMigrate(&series.Series{})
	database = db
}

func GetDatabase() (DB *gorm.DB) {
	return database
}

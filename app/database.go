package app

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + host +")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
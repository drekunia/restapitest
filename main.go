package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB
var err error

func main() {
	dbHost := os.Getenv("APP_DB_HOST")
	dbName := os.Getenv("APP_DB_NAME")
	dbPort := os.Getenv("APP_DB_PORT")
	dbUsername := os.Getenv("APP_DB_USERNAME")
	dbPassword := os.Getenv("APP_DB_PASSWORD")

	dbUri := fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s sslmode=disable", dbHost, dbName, dbPort, dbUsername, dbPassword)

	db, err = gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db.Debug().AutoMigrate(&User{})
	handleRequest()
}

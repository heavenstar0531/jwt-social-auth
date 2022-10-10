package database

import (
	"fmt"
	"jwt-go/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	dbPort   = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "jwt-go"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbname)
	dsn := config

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connecting to Database: ", err)
	}

	fmt.Println("Database Connected Successfully")
	db.Debug().AutoMigrate(&models.User{}, &models.Product{})
}

func GetDB() *gorm.DB {
	return db
}

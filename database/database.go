package database

import (
	models2 "github.com/decadevs/next_store/models"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

func DB() *gorm.DB {

	username := os.Getenv("root")
	password := os.Getenv("OluwaTimi30")
	dbname := os.Getenv("e-commerce_db")

	//Database connection
	db, err := gorm.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	AutoMigrate(db)
	return db
}

//GORM AUTO-MIGRATION OF DATABASE
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models2.User{},
		&models2.Buyer{},
		&models2.Order{},
		&models2.Product{},
		&models2.Status{},
		&models2.Seller{})
}

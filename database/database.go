package database

import (
	"fmt"
	"github.com/decadevs/next_store/models"
	_ "github.com/decadevs/next_store/models"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func DB() *gorm.DB {

	const username = ""
	const password = ""
	const dbname = ""

	//Database connection
	db, err := gorm.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local")
	//error handling
	if err != nil {
		log.Println("checking database error", err)
	}
	defer db.Close()

	AutoMigrate(db)
	log.Println("got here")
	return db

}

//GORM AUTO-MIGRATION OF DATABASE
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{},
		&models.Buyer{},
		&models.Order{},
		&models.Product{},
		&models.Status{},
		&models.Seller{})

	log.Println("checking database error", err)
}

//CREATE A IN-MEMORY SELLER DATABASE-RECORD
func SellerDB() *gorm.DB {

	//open database connection
	const username = ""
	const password = ""
	const dbname = ""

	//Database connection
	db, err := gorm.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local")
	//error handling
	if err != nil {
		log.Println("checking database error", err)
	}
	defer db.Close()

	currentTime := time.Now()
	Model := gorm.Model{ID: 1, CreatedAt: currentTime, UpdatedAt: currentTime, DeletedAt: nil}
	var Product = []models.Product{{Model, "Golang Squard10 Polo", 6000, 10, "Nice Polo Outfit", "https://www.google.com/url?sa=i&url=https%3A%2F%2Fjiji.ng%2Fajah%2Fclothing%2Ft-shirt-and-cap-print-branding-oVVMEnC1mBFSzgFoMhPoYd5R.html&psig=AOvVaw19zEvHw76sOOwj0T8hLpRO&ust=1650138377657000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCIi5ttnqlvcCFQAAAAAdAAAAABAD"}}

	//CREATE
	var Seller = models.Seller{
		models.User{1, "Golang-SQ10-Ecommerce", "sq10golang@gmail.com", "ProductOwner", "12345", "12345", "Edo Tech Park"}, 1, 1,
		Product}

	result := db.Create(&Seller)
	err = result.Error
	if err != nil {
		log.Printf("Error creating seller record : %v", err)
	}
	rowsAffected := result.RowsAffected
	fmt.Println("Number of rows affected", rowsAffected)

	return db
}

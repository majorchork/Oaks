package main

import (
	models2 "github.com/decadevs/next_store/models"
	"github.com/decadevs/next_store/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/driver/mysql"
	"log"
	_ "net/http"
)

func main() {

	//call routes /sever
	routes.CallRoutes("port")

	//Database connection
	db, err := gorm.Open("mysql", "root:OluwaTimi30@tcp(127.0.0.1:3306)/e-commerce_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	AutoMigrate(db)

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

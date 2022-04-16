package main

import (
	"github.com/decadevs/next_store/database"
	_ "github.com/decadevs/next_store/models"
	"github.com/decadevs/next_store/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	_ "gorm.io/driver/mysql"
	_ "net/http"
)

func main() {
	//database connection
	database.DB()

	//seller's in-memory data
	database.SellerDB()

	//call routes /sever
	routes.CallRoutes("port")

}

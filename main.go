package main

import (
	"github.com/decadevs/next_store/database"
	"github.com/decadevs/next_store/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/driver/mysql"
	_ "net/http"
)

func main() {

	//call routes /sever
	routes.CallRoutes("port")

	database.DB()

}

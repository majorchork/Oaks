package routes

import (
	"github.com/decadevs/next_store/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

//THE ROUTING IS USED TO HANDLE VARIOUS URL
func CallRoutes(port string, db *gorm.DB) {
	//set route as default one made by Gin
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	//serve the static files
	router.StaticFS("static", http.Dir("./templates/static"))

	//sever all the HTML template quickly as soon as the pages load
	router.LoadHTMLGlob("templates/*.html")

	//Welcome page router
	router.GET("/", handlers.Welcomepage)

	//Market place router
	router.GET("/marketplace", handlers.MarketPlace)

	//buyer page router
	router.GET("/addtocart", handlers.AddToCart)

	//Buyer SignUp router
	router.GET("/buyerssignup", handlers.BuyerSignUp)

	//Seller Edit Product
	router.GET("/sellers/editPost/:id", handlers.SellerEditProduct)

	//Seller Update Product
	router.POST("/update-product/:id", handlers.SellerUpdateProduct)

	//Admin Post Product
	router.POST("/sellers/addproducts", handlers.AdminPostProduct)

	//Admin Get Product
	router.GET("/sellers/addproductspage", handlers.AdminGetProduct)

	//Admin Delete Product
	router.GET("/sellers/deleteproduct/:id", handlers.AdminDeleteProduct)

	//Admin Launch Product to Market Place
	router.POST("/seller/postproduct", handlers.AdminPostInMarket)
	router.GET("/sellers/launchproduct", handlers.AdminLaunchProduct)

	//SIGN UP AND LOGIN
	//seller signin router
	router.GET("/sellersignup", handlers.SellerLogin)

	//Seller page router
	router.GET("/sellerpage", handlers.SellerPage)

	//Seller Login
	router.POST("/sellers/signin", handlers.SellerLoginHandler)

	//Buyer Sig-up
	router.POST("/buyers/signup", handlers.BuyerSignUpHandler)
	//Buyer Login-in
	router.POST("/buyers/login", handlers.LoginHandler)

	//To search for product
	router.GET("/searchproduct", handlers.SearchProduct)

	//Admin get notified once a buyer makes an order
	router.POST("/orderpage", handlers.PayNow)

	//start and run the server on port 8084
	port = ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8084"
	}
	router.Run(port)
}

package handlers

import (
	"github.com/decadevs/next_store/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
)

//external database
var db *gorm.DB

//HANDLE WELCOME PAGE
func Welcomepage(c *gin.Context) {

	//call the HTML Method of the context to render the template
	c.HTML(
		//setup the status of the template
		http.StatusOK,

		//use which template or deploy which template
		"index.html",
		gin.H{
			"title": "Next-Store"})

}

//HANDLER MARKET PLACE
func MarketPlace(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"marketplace.html",
		gin.H{
			"message": "Market Place"})
}

//HANDLER BUYER SIGN UP PAGE
func BuyerSignUp(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"buyer_signup.html",
		gin.H{
			"message": "Buyer Sign In"})
}

//HANDLER BUYER PAGE
func BuyerPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"buyer_page.html",
		gin.H{
			"message": "Buyer Page"})
}

//HANDLER SELLER SIGN UP PAGE
func SellerSignUp(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"seller_signup.html",
		gin.H{
			"message": "Seller Sign Up"})
}

//HANDLER SELLER PAGE
func SellerPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"seller_page.html",
		gin.H{
			"message": "Seller Page"})
}

func AddProductToMarket(c *gin.Context) {

	//initialize a variable of type product
	var product models.Product

	//parsing the form values
	product.Image = c.PostForm("product-image")
	product.Name = c.PostForm("product-Name")
	product.Quantity, _ = strconv.Atoi(c.PostForm("product-quantity-left"))
	product.Price, _ = strconv.Atoi(c.PostForm("product-price"))

	//initialize the database

	err := db.Save(product).Error
	if err != nil {
		log.Println("Failed to save product in database")
	}
	c.Redirect(http.StatusMovedPermanently, "/marketplace")
}

func EditProduct(c *gin.Context) {

}

func UpdateProduct(c *gin.Context) {

}

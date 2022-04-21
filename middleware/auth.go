package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Authentication(c *gin.Context) *http.Cookie {
	email, err := c.Request.Cookie("email")
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusPermanentRedirect, "/sellersignup")
		return nil
	}
	return email
}

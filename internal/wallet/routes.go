package wallet

import (
	"algatux/wallet/internal/authenticator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var auth = authenticator.Auth{}

var getPing gin.HandlerFunc = func(c *gin.Context) {

	claims, _ := c.Keys["claims"].(*jwt.RegisteredClaims)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"subject": claims.Subject,
	})
}

var postLogin gin.HandlerFunc = func(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	token, err := auth.AuthenticateCredentials(username, password)

	if nil != err {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

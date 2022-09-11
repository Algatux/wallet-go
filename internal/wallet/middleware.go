package wallet

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

var authMiddleware gin.HandlerFunc = func(c *gin.Context) {

	if c.Request.URL.Path == "/auth/login" {
		return
	}

	authToken := c.Request.Header.Get("Authorization")
	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	if false == strings.HasPrefix(authToken, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	parts := strings.Split(authToken, " ")
	err := auth.ValidateToken(parts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}
}

var parseAuthToken gin.HandlerFunc = func(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	authToken = strings.ReplaceAll(authToken, "Bearer ", "")
	parts := strings.Split(authToken, ".")
	uDec, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		panic(err)
	}

	data := &jwt.RegisteredClaims{}
	err = json.Unmarshal(uDec, &data)
	if err != nil {
		panic(err)
	}

	c.Keys = make(map[string]any)
	c.Keys["claims"] = data
}

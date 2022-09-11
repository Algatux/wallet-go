package wallet

import (
	"algatux/wallet/internal/authenticator"
	"github.com/gin-gonic/gin"
)

type Wallet struct {
	router *gin.Engine
	auth   *authenticator.Auth
}

func (w *Wallet) Boot(pair authenticator.RsaKeyPair) *Wallet {
	w.auth = &auth
	w.auth.SetRsaKeyPair(pair)
	w.router = gin.Default()
	w.registerMiddlewares()
	w.registerRoutes()

	return w
}

func (w *Wallet) Start(address string) {
	err := w.router.Run(address)
	if err != nil {
		panic(err)
	}
}

func (w *Wallet) registerRoutes() {
	w.router.GET("/ping", getPing)
	group := w.router.Group("/auth")
	group.POST("/login", postLogin)
}

func (w *Wallet) registerMiddlewares() {
	w.router.Use(gin.Recovery())
	w.router.Use(authMiddleware)
	w.router.Use(parseAuthToken)
}

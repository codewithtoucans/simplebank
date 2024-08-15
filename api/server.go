package api

import (
	db "github.com/codewithtoucans/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *server {
	server := &server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *server) Start(address string) error {
	return server.router.Run(address)
}

func errReponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

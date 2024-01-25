package api

import (
	"net/http"

	db "github.com/codewithtoucans/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	server.router = router

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("accounts", server.createAccount)
	router.GET("accounts/:id", server.getAccount)
	router.GET("accounts", server.listAccount)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

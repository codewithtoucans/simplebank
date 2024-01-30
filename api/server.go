package api

import (
	"fmt"
	"net/http"

	"github.com/codewithtoucans/simplebank/config"
	db "github.com/codewithtoucans/simplebank/db/sqlc"
	"github.com/codewithtoucans/simplebank/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     config.Config
	tokenMaker token.Maker
	store      *db.Store
	router     *gin.Engine
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: token,
	}
	router := gin.Default()
	server.router = router

	_ = binding.Validator.Engine().(*validator.Validate).RegisterValidation("currency", validCurrency)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("accounts", server.createAccount)
	authRoutes.GET("accounts/:id", server.getAccount)
	authRoutes.GET("accounts", server.listAccount)
	authRoutes.POST("/transfers", server.createTransfer)

	return server, nil
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

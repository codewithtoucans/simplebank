package api

import (
	db "github.com/codewithtoucans/simplebank/db/sqlc"
	"github.com/codewithtoucans/simplebank/token"
	"github.com/codewithtoucans/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type server struct {
	config     util.Config
	tokenMaker token.Maker
	store      db.Store
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) *server {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil
	}

	server := &server{store: store, tokenMaker: tokenMaker, config: config}

	server.setupRouter()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	return server
}

func (server *server) setupRouter() {
	router := gin.Default()
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfer", server.createTransfer)
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	server.router = router
}

func (server *server) Start(address string) error {
	return server.router.Run(address)
}

func errReponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

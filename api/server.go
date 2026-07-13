package api

import (
	"fmt"

	db "github.com/Nevets9/simplebank/db/sqlc"
	"github.com/Nevets9/simplebank/token"
	"github.com/Nevets9/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP request for banking service
type Server struct {
	config util.Config
	store db.Store
	router *gin.Engine
	tokenMaker token.Maker
}

// NewServer create a new HTTP server and setup routing 
func NewServer(config util.Config, store db.Store) (*Server, error) {
	//tokenMaker, err := token.NewPasetoMaker(config.TokenSecret)
	tokenMaker, err := token.NewJWTMaker(config.TokenSecret) 
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker:%w", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter(){
	router := gin.Default()

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a spesific address.
func (server *Server) Start (address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error ": err.Error()}
}
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/salman1s2h/simplebank/db/sqlc"
	"github.com/salman1s2h/simplebank/token"
	"github.com/salman1s2h/simplebank/util"
)

type Server struct {
	config     *util.Env
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config *util.Env, Store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TOKEN_SYMMETRIC_KEY)
	if err != nil {
		panic(err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		store:      Store,
	}

	// router := gin.Default()
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("failed to get validator engine")
	}

	if err := v.RegisterValidation("currency", validCurrency); err != nil {
		panic(err)
	}

	// router.POST("/accounts", server.createAccount)
	// router.GET("/accounts/:id", server.getAccount)
	// router.GET("/accounts", server.listAccounts)
	// router.POST("/users", server.createUser)
	// router.GET("/users/:username", server.getUser)

	// router.POST("/transfers", server.createTransfer)
	// server.router = router
	server.setupRouter()
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	// router.POST("/users/login", server.loginUser)
	// router.POST("/tokens/renew_access", server.renewAccessToken)

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on the given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

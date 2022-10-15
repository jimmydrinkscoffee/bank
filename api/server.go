package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vbph/bank/db/sqlc"
	"github.com/vbph/bank/token"
	"github.com/vbph/bank/utils"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     utils.Config
}

func CreateServer(config utils.Config, dbStore *db.Store) (*Server, error) {
	tokenMaker, err := token.CreatePasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		store:      dbStore,
		tokenMaker: tokenMaker,
		config:     config,
	}

	server.initRouter()

	return server, nil
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func (server *Server) initRouter() {
	router := gin.Default()

	router.POST("/auth/sign-up", server.signUp)
	router.POST("/auth/login", server.login)

	router.GET("/account/:id", server.readAccount)
	router.DELETE("/account", server.deleteAccount)

	router.POST("/transfer", server.transfer)

	server.router = router
}

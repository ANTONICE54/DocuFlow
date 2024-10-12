package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Verify(ctx *gin.Context)
}

type Server struct {
	router      *gin.Engine
	userHandler IUserHandler
}

func New(userH IUserHandler) *Server {
	s := Server{
		router:      gin.Default(),
		userHandler: userH,
	}
	s.setUpRoutes()

	return &s
}

func (s *Server) setUpRoutes() {

	s.router.POST("/register", s.userHandler.Register)
	s.router.POST("/login", s.userHandler.Login)
	s.router.POST("/verify", s.userHandler.Verify)
}

func (s *Server) Run(serverPort string) {
	if err := s.router.Run(":" + serverPort); err != nil {
		log.Fatalf("Failed to run server %v", err.Error())
	}
}

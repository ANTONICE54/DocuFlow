package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ICategoryHandler interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ISubcategoryHandler interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Server struct {
	router             *gin.Engine
	categoryHandler    ICategoryHandler
	subcategoryHandler ISubcategoryHandler
}

func New(categoryH ICategoryHandler, subcategoryH ISubcategoryHandler) *Server {
	s := Server{
		router:             gin.Default(),
		categoryHandler:    categoryH,
		subcategoryHandler: subcategoryH,
	}
	s.setUpRoutes()

	return &s
}

func (s *Server) setUpRoutes() {

	s.router.POST("/category", s.categoryHandler.Create)
	s.router.GET("/category", s.categoryHandler.List)
	s.router.PATCH("/category/:id", s.categoryHandler.Update)
	s.router.DELETE("/category/:id", s.categoryHandler.Delete)

	s.router.POST("/subcategory", s.subcategoryHandler.Create)
	s.router.GET("/subcategory", s.subcategoryHandler.List)
	s.router.PATCH("/subcategory/:id", s.subcategoryHandler.Update)
	s.router.DELETE("/subcategory/:id", s.subcategoryHandler.Delete)
}

func (s *Server) Run(serverPort string) {
	if err := s.router.Run(":" + serverPort); err != nil {
		log.Fatalf("Failed to run server %v", err.Error())
	}
}

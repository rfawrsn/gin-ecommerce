package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	api := r.Group("/api")
	HelloWorldRoute(api)

	return r
}

func HelloWorldRoute(r *gin.RouterGroup) {
	r.GET("/hello", HelloWorldController)
}

func HelloWorldController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from Task Manager"})
}

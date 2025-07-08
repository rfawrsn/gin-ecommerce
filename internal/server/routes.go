package server

import (
	"net/http"

	"github.com/fawwazalifiofarsa/gin-ecommerce/internal/controllers"
	"github.com/fawwazalifiofarsa/gin-ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	api := r.Group("/api")
	HelloWorldRoute(api)
	ProductRoutes(api)
	ProtectedRoutes(api)

	return r
}

func HelloWorldRoute(r *gin.RouterGroup) {
	r.GET("/hello", HelloWorldController)
}

func ProtectedRoutes(r *gin.RouterGroup) {
	// Apply auth middleware to all routes in this group
	protected := r.Group("/protected")
	protected.Use(middlewares.AuthMiddleware())

	// Test protected endpoint
	protected.GET("/test", ProtectedTestController)
}

func ProductRoutes(r *gin.RouterGroup) {
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
}

func HelloWorldController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from Task Manager"})
}

func ProtectedTestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Access granted! This is a protected endpoint",
		"status":  "authenticated",
	})
}

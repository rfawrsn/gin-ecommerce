package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fawwazalifiofarsa/gin-ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// POST /products
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.InsertProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product inserted successfully"})
}

// GET /products/:name
func GetProductByName(c *gin.Context) {
	name := c.Param("name")
	product, err := models.Find(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// GET /products
func GetAllProducts(c *gin.Context) {
	name := c.Query("name")
	if name != "" {
		product, err := models.Find(name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
		return
	}
	products := []models.Product{}
	collection := models.MongoClient.Database(models.DB).Collection(models.CollName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("DB Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		fmt.Println("DB Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// PUT /products/:id
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateProduct(id, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DELETE /products/:id
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// GET /products/:id
func GetProductById(c *gin.Context) {
	id := c.Param("id")
	product, err := models.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

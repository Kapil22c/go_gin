package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Kapil22c/go_gin/bookstore_RESTAPI/models"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Kapil22c/go_gin/bookstore_RESTAPI/controllers"
	"github.com/Kapil22c/go_gin/bookstore_RESTAPI/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run()
}

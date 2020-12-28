package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "github.com/Kapil22c/go_gin/assignment-7/CRUD_Gin_postgresql/config"
	routes "github.com/Kapil22c/go_gin/assignment-7/CRUD_Gin_postgresql/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}

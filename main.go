package main

import (
	"LibraryManagement/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app := gin.Default()

	router := app.Group("/api/v1")
	routes.AddRoutes(router)

	err := app.Run(":8080")
	if err != nil {
		log.Println("Failed to start server")
		return
	}
}

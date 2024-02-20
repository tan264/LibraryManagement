package main

import (
	"LibraryManagement/controller"
	"LibraryManagement/database"
	"LibraryManagement/middleware"
	"LibraryManagement/repository"
	"LibraryManagement/routes"
	"LibraryManagement/service"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	db = database.Connect()

	userRepository = repository.UserRepository(db)

	authService = service.AuthService(userRepository)
	jwtService  = service.NewJwtService()
	userService = service.NewUserService(userRepository)

	authController = controller.AuthController(authService, jwtService)
	userController = controller.NewUserController(userService)
)

func main() {
	router := gin.Default()

	authRoutes := router.Group("/auth")
	routes.AuthRoutes(authRoutes, authController)

	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.AuthorizeJWT(jwtService))
	routes.UserRoutes(userRoutes, userController)

	err := router.Run(":8080")
	if err != nil {
		log.Println("Failed to start server")
		return
	}
}

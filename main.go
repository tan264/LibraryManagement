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
	bookRepository = repository.NewBookRepository(db)

	authService  = service.AuthService(userRepository)
	jwtService   = service.NewJwtService()
	userService  = service.NewUserService(userRepository)
	adminService = service.NewAdminService(bookRepository)

	authController  = controller.AuthController(authService, jwtService)
	userController  = controller.NewUserController(userService)
	adminController = controller.NewAdminController(adminService)
)

func main() {
	router := gin.Default()

	authRoutes := router.Group("/auth")
	routes.AuthRoutes(authRoutes, authController)

	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.AuthorizeJWT(jwtService))
	routes.UserRoutes(userRoutes, userController)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthorizeJWT(jwtService))
	routes.AdminRoutes(adminRoutes, adminController)

	err := router.Run(":8080")
	if err != nil {
		log.Println("Failed to start server")
		return
	}
}

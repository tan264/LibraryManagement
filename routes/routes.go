package routes

import (
	"LibraryManagement/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(authRoutes *gin.RouterGroup, authController controller.IAuthController) {
	authRoutes.POST("/register", authController.Register)
	authRoutes.POST("/login", authController.Login)
}

func UserRoutes(authRoutes *gin.RouterGroup, userController controller.IUserController) {
	authRoutes.PUT("/edit", userController.EditAccount)
	authRoutes.DELETE("/delete", userController.DeleteAccount)
	authRoutes.GET("/filter", userController.FilterAccount)
	authRoutes.GET("/report-by-created_at", userController.StatisticizeAccountByCreatedAt)
}

func AdminRoutes(adminRoutes *gin.RouterGroup, adminController controller.IAdminController) {
	adminRoutes.POST("/import", adminController.ImportBookData)
}

package routes

import (
	"LibraryManagement/controller"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func AuthRoutes(authRoutes *gin.RouterGroup, authController controller.IAuthController) {
	authRoutes.POST("/register", authController.Register)
	authRoutes.POST("/login", authController.Login)
}

func UserRoutes(userRoutes *gin.RouterGroup, userController controller.IUserController) {
	userRoutes.PUT("/edit", userController.EditAccount)
	userRoutes.DELETE("/delete", userController.DeleteAccount)
	userRoutes.GET("/filter", userController.FilterAccount)
	userRoutes.GET("/report-by-created_at", userController.StatisticizeAccountByCreatedAt)
	userRoutes.POST("/checkout/:book_id", userController.CheckoutBook)
	userRoutes.POST("/return/:book_id", userController.ReturnBook)
}

func AdminRoutes(adminRoutes *gin.RouterGroup, adminController controller.IAdminController) {
	adminRoutes.POST("/import", adminController.ImportBookData)
	adminRoutes.PUT("/update", adminController.UpdateBook)
	adminRoutes.POST("/create", adminController.CreateBook)
	adminRoutes.DELETE("/delete/:book_id", adminController.DeleteBook)
	adminRoutes.GET("/get-pdf-report", func(context *gin.Context) {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "Hello, world")
		pdf.OutputFileAndClose("hello.pdf")
		context.File("hello.pdf")
	})
}

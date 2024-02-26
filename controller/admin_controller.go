package controller

import (
	"LibraryManagement/service"
	"github.com/gin-gonic/gin"
)

type IAdminController interface {
	ImportBookData(context *gin.Context)
}

type adminController struct {
	adminService service.IAdminService
}

func NewAdminController(adminService service.IAdminService) IAdminController {
	return adminController{
		adminService: adminService,
	}
}

func (a adminController) ImportBookData(context *gin.Context) {
	file, _ := context.FormFile("file")
	message, err := a.adminService.ImportBookData(file)
	if err != nil {
		context.JSON(500, gin.H{
			"message": message,
		})
	} else {
		context.JSON(200, gin.H{
			"message": message,
		})
	}
}

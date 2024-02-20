package controller

import (
	"LibraryManagement/model"
	"LibraryManagement/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserController interface {
	EditAccount(context *gin.Context)
	DeleteAccount(context *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func (u UserController) EditAccount(context *gin.Context) {
	var editRequest model.EditAccountRequest
	err := context.ShouldBindJSON(&editRequest)
	if err != nil {
		response := model.BuildResponse("Failed to process request", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userID, exists := context.Get("userID")
	if exists {
		userUpdated, err := u.userService.EditAccount(editRequest, uint(userID.(float64)))
		if err != nil {
			response := model.BuildResponse("Failed to process request", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		response := model.BuildResponse("Update success", "", userUpdated)
		context.JSON(http.StatusOK, response)
	} else {
		response := model.BuildResponse("Failed to process request", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (u UserController) DeleteAccount(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewUserController(userService service.IUserService) UserController {
	return UserController{
		userService: userService,
	}
}

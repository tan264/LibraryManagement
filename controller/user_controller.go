package controller

import (
	"LibraryManagement/model"
	"LibraryManagement/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUserController interface {
	EditAccount(context *gin.Context)
	DeleteAccount(context *gin.Context)
	FilterAccount(context *gin.Context)
	StatisticizeAccountByCreatedAt(context *gin.Context)
	CheckoutBook(context *gin.Context)
	ReturnBook(context *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func (u UserController) CheckoutBook(context *gin.Context) {
	userID, exists := context.Get("userID")
	if exists {
		bookIDString := context.Param("book_id")
		bookID, err := strconv.ParseUint(bookIDString, 10, 64)
		if err != nil {
			response := model.BuildResponse("Failed to process request", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		checkout, err := u.userService.CheckoutBook(uint(bookID), uint(userID.(float64)))
		if err != nil {
			response := model.BuildResponse("Failed to checkout", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := model.BuildResponse("Checkout success", "", checkout)
		context.JSON(http.StatusOK, response)
	} else {
		response := model.BuildResponse("Unauthorized", "", nil)
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}

func (u UserController) ReturnBook(context *gin.Context) {
	userID, exists := context.Get("userID")
	if exists {
		bookIDString := context.Param("book_id")
		bookID, err := strconv.ParseUint(bookIDString, 10, 64)
		if err != nil {
			response := model.BuildResponse("Failed to process request", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		checkout, err := u.userService.ReturnBook(uint(bookID), uint(userID.(float64)))
		if err != nil {
			response := model.BuildResponse("Failed to return", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := model.BuildResponse("Return success", "", checkout)
		context.JSON(http.StatusOK, response)
	} else {
		response := model.BuildResponse("Unauthorized", "", nil)
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)

	}
}

func (u UserController) StatisticizeAccountByCreatedAt(context *gin.Context) {
	result, err := u.userService.StatisticizeAccountByCreatedAt()
	if err != nil {
		response := model.BuildResponse("Failed to filter account", "", nil)
		context.JSON(http.StatusInternalServerError, response)
	}
	response := model.BuildResponse("Success", "", result)
	context.JSON(http.StatusOK, response)
}

func (u UserController) FilterAccount(context *gin.Context) {
	var userToFilter model.FilterUserRequest
	err := context.ShouldBindQuery(&userToFilter)
	if err != nil {
		response := model.BuildResponse("Failed to process request", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	users, err := u.userService.FilterAccount(userToFilter)
	if err != nil {
		response := model.BuildResponse("Failed to filter account", "", nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	response := model.BuildResponse("Filter Success", "", users)
	context.JSON(http.StatusOK, response)
}

func NewUserController(userService service.IUserService) IUserController {
	return UserController{
		userService: userService,
	}
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
			response := model.BuildResponse("Failed to update", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		response := model.BuildResponse("Update success", "", userUpdated)
		context.JSON(http.StatusOK, response)
	} else {
		response := model.BuildResponse("Unauthorized", "", nil)
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}

func (u UserController) DeleteAccount(context *gin.Context) {
	userID, exists := context.Get("userID")
	if exists {
		err := u.userService.DeleteAccount(uint(userID.(float64)))
		if err != nil {
			response := model.BuildResponse("Failed to delete", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		response := model.BuildResponse("Delete success", "", userID)
		context.JSON(http.StatusOK, response)
	} else {
		response := model.BuildResponse("Unauthorized", "", nil)
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}

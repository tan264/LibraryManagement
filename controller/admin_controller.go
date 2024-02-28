package controller

import (
	"LibraryManagement/model"
	"LibraryManagement/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IAdminController interface {
	ImportBookData(context *gin.Context)
	CreateBook(context *gin.Context)
	UpdateBook(context *gin.Context)
	DeleteBook(context *gin.Context)
}

type adminController struct {
	bookService service.IBookService
}

func (a adminController) CreateBook(context *gin.Context) {
	var input model.CreateBookRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		response := model.BuildResponse("Failed to process request", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	bookCreated, err := a.bookService.CreateBook(input)
	if err != nil {
		response := model.BuildResponse("Failed to create book", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
	} else {
		response := model.BuildResponse("Success", "", bookCreated)
		context.JSON(http.StatusOK, response)

	}
}

func (a adminController) UpdateBook(context *gin.Context) {
	var input model.UpdateBookRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		response := model.BuildResponse("Failed to process request", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	bookUpdated, err := a.bookService.UpdateBook(input)
	if err != nil {
		response := model.BuildResponse("Failed to update book", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
	} else {
		response := model.BuildResponse("Success", "", bookUpdated)
		context.JSON(http.StatusOK, response)
	}
}

func (a adminController) DeleteBook(context *gin.Context) {
	userIDString := context.Param("book_id")
	userID, err := strconv.ParseUint(userIDString, 10, 64)
	if err != nil {
		response := model.BuildResponse("Invalid input", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	err = a.bookService.DeleteBook(uint(userID))
	if err != nil {
		response := model.BuildResponse("Failed to delete book", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	response := model.BuildResponse("Success", "", nil)
	context.JSON(http.StatusOK, response)
}

func NewAdminController(bookService service.IBookService) IAdminController {
	return adminController{
		bookService: bookService,
	}
}

func (a adminController) ImportBookData(context *gin.Context) {
	file, _ := context.FormFile("file")
	message, err := a.bookService.ImportBookData(file)
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

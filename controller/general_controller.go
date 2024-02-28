package controller

import (
	"LibraryManagement/model"
	"LibraryManagement/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IGeneralController interface {
	GetLibrariesOfAddress(context *gin.Context)
	GetBooksOfLibrary(context *gin.Context)
}

type generalController struct {
	addressService service.IAddressService
	libraryService service.ILibraryService
}

func (g generalController) GetLibrariesOfAddress(context *gin.Context) {
	addressName := context.Param("address_name")
	libraries, err := g.addressService.GetLibraries(addressName)
	if err != nil {
		response := model.BuildResponse("Failed to get libraries", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	response := model.BuildResponse("Success", "", libraries)
	context.JSON(http.StatusOK, response)
}

func (g generalController) GetBooksOfLibrary(context *gin.Context) {
	libraryIDString := context.Param("library_id")
	libraryID, err := strconv.ParseUint(libraryIDString, 10, 64)
	if err != nil {
		response := model.BuildResponse("Invalid input", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	books, err := g.libraryService.GetBooks(uint(libraryID))
	if err != nil {
		response := model.BuildResponse("Failed to get books", err.Error(), nil)
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	response := model.BuildResponse("Success", "", books)
	context.JSON(http.StatusOK, response)
}

func NewGeneralController(addressService service.IAddressService, libraryService service.ILibraryService) IGeneralController {
	return generalController{
		addressService: addressService,
		libraryService: libraryService,
	}
}

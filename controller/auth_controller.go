package controller

import (
	"LibraryManagement/model"
	"LibraryManagement/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type authController struct {
	authService service.IAuthService
	jwtService  service.IJwtService
}

func (a authController) Login(context *gin.Context) {
	var loginRequest model.LoginRequest
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		response := model.BuildResponse("Failed to process request", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	loggedUserID, err := a.authService.VerifyCredential(loginRequest.Username, loginRequest.Password)
	if err != nil {
		response := model.BuildResponse("Login Failed!", "Invalid username or password", nil)
		context.JSON(http.StatusConflict, response)
	} else {
		token, err := a.jwtService.GenerateToken(loggedUserID)
		if err != nil {
			response := model.BuildResponse("Something wrong", "Internal error", nil)
			context.JSON(http.StatusInternalServerError, response)
		} else {
			response := model.BuildResponse("Login success", "", token)
			context.JSON(http.StatusOK, response)
		}
	}
}

func (a authController) Register(context *gin.Context) {
	var input model.RegisterUserRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		response := model.BuildResponse("Failed to process request", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if a.authService.IsDuplicateUsername(input.Username) {
		response := model.BuildResponse("Failed to register new user", "Username already exists", input)
		context.JSON(http.StatusConflict, response)
	} else {
		registeredUser := a.authService.RegisterUser(input)
		response := model.BuildResponse("Register success", "", registeredUser)
		context.JSON(http.StatusOK, response)
	}
}

func AuthController(authService service.IAuthService, jwtService service.IJwtService) IAuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

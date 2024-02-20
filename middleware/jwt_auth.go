package middleware

import (
	"LibraryManagement/model"
	"LibraryManagement/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := model.BuildResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader[7:])
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("userID", claims["user_id"])
			c.Next()
		} else {
			response := model.BuildResponse("token invalid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}

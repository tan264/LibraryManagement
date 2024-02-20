package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type IJwtService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtService struct {
	secretKey string
}

func (jwtService JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(jwtService.secretKey), nil
	})
	return token, err
}

func NewJwtService() JwtService {
	secretKey := os.Getenv("SECRET_KEY")
	return JwtService{
		secretKey,
	}
}

func (jwtService JwtService) GenerateToken(userID uint) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	tokenString, err = token.SignedString([]byte(jwtService.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

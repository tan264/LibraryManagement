package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type RegisterInput struct {
	UserID         uint           `json:"userId" gorm:"primary_key; auto_increment"`
	Username       string         `json:"username" gorm:"not_null; unique" binding:"required"`
	Password       string         `json:"password" gorm:"not_null" binding:"required"`
	Email          sql.NullString `json:"email"`
	Phone          sql.NullString `json:"phone"`
	FirstName      sql.NullString `json:"firstName"`
	LastName       sql.NullString `json:"lastName"`
	DateRegistered time.Time      `json:"dateRegistered" gorm:"not_null; default:current_timestamp()"`
	IsActive       uint           `json:"isActive" gorm:"type:tinyint(1); default:1"`
	IsAdmin        uint           `json:"isAdmin" gorm:"type:tinyint(1); default:0"`
}

func Register(context *gin.Context, db *gorm.DB) {
	var input RegisterInput
	
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": input})
		return
	}
	
	hashedPassword, err := HashPassword(input.Password)
	input.Password = hashedPassword
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": input})
		return
	}
	
	db.Create(&input)
	
	context.JSON(http.StatusOK, gin.H{"message": "Register successfully", "data": input})
	
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

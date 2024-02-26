package model

import (
	"time"
)

type User struct {
	UserID    uint       `json:"user_id" gorm:"primary_key; auto_increment"`
	Username  string     `json:"username" gorm:"type:varchar(40); not null; unique" binding:"required"`
	Password  string     `json:"-" gorm:"not null" binding:"required"`
	Email     *string    `json:"email" gorm:"type:varchar(40)"`
	Phone     *string    `json:"phone" gorm:"type:varchar(20)"`
	FirstName *string    `json:"first_name" gorm:"type:varchar(40)"`
	LastName  *string    `json:"last_name" gorm:"type:varchar(40)"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:date not null"`
	IsActive  uint8      `json:"is_active" gorm:"type:tinyint(1);not null; default:1"`
	IsAdmin   uint8      `json:"is_admin" gorm:"type:tinyint(1);not null; default:0"`
	Checkouts []Checkout `json:"checkouts"`
	Tokens    []Token    `json:"tokens"`
}

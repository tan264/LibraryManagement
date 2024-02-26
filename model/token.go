package model

import "time"

type Token struct {
	ID        uint      `json:"id"  gorm:"primary_key; auto_increment"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"type:varchar(255); not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime; not null"`
	Platform  string    `json:"platform" gorm:"type:varchar(40)"`
}

package model

import "time"

type Checkout struct {
	ID         uint      `json:"id" gorm:"primary_key; auto_increment"`
	StartTime  time.Time `json:"start_time" gorm:"type:datetime; not null"`
	EndTime    time.Time `json:"end_time" gorm:"type:datetime; not null"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	BookID     uint      `json:"book_id" gorm:"not null"`
	IsReturned uint8     `json:"is_returned" gorm:"type:tinyint(1); not null; default: 0"`
}

func (checkout Checkout) ToString() string {
	return checkout.StartTime.Format("2006-01-02 15:04:05") + " - " + checkout.EndTime.Format("2006-01-02 15:04:05")
}

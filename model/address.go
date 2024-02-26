package model

type Address struct {
	AddressID uint      `json:"address_id" gorm:"primary_key; auto_increment"`
	Name      string    `json:"name" gorm:"type:varchar(255); not null"`
	Libraries []Library `json:"libraries"`
}

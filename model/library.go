package model

type Library struct {
	LibraryID   uint   `json:"library_id" gorm:"primary_key; auto_increment"`
	LibraryName string `json:"name" gorm:"type:varchar(255); not null"`
	AddressID   uint   `json:"address_id" gorm:"not null"`
	Books       []Book `json:"books"`
}

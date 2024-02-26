package model

type Book struct {
	BookID    uint       `json:"book_id" gorm:"primary_key; auto_increment"`
	Title     string     `json:"title" gorm:"type:varchar(255); not null"`
	LibraryID uint       `json:"library_id" gorm:"not null"`
	Checkouts []Checkout `json:"checkouts"`
}

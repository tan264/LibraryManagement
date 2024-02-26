package repository

import (
	"LibraryManagement/model"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(book model.Book) (model.Book, error)
}

type bookRepository struct {
	dbConnection *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return bookRepository{
		dbConnection: db,
	}
}

func (b bookRepository) CreateBook(book model.Book) (model.Book, error) {
	err := b.dbConnection.Create(&book).Error
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

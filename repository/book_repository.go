package repository

import (
	"LibraryManagement/model"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(book model.Book) (model.Book, error)
	UpdateBook(book model.Book) (model.Book, error)
	DeleteBook(bookID uint) error
	GetBookByID(bookID uint) (model.Book, error)
	FilterBook(bookToFilter model.FilterBookRequest) ([]model.Book, error)
}

type bookRepository struct {
	dbConnection *gorm.DB
}

func (b bookRepository) FilterBook(bookToFilter model.FilterBookRequest) ([]model.Book, error) {
	var books []model.Book
	query := b.dbConnection.Find(&books)
	if bookToFilter.BookID > 0 {
		query = query.Where("book_id = ?", bookToFilter.BookID)
	}
	if bookToFilter.CreatedAt != "" {
		query = query.Where("DATE(created_at) = ?", bookToFilter.CreatedAt)
	}
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b bookRepository) UpdateBook(book model.Book) (model.Book, error) {
	err := b.dbConnection.Save(&book).Error
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (b bookRepository) DeleteBook(bookID uint) error {
	err := b.dbConnection.Delete(&model.Book{}, bookID).Error
	if err != nil {
		return err
	}
	return nil
}

func (b bookRepository) GetBookByID(bookID uint) (model.Book, error) {
	var book model.Book
	err := b.dbConnection.First(&book, bookID).Error
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
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

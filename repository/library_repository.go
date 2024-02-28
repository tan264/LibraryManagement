package repository

import (
	"LibraryManagement/model"
	"gorm.io/gorm"
)

type ILibraryRepository interface {
	GetBooks(libraryID uint) ([]model.Book, error)
}

type libraryRepository struct {
	dbConnection *gorm.DB
}

func (l libraryRepository) GetBooks(libraryID uint) ([]model.Book, error) {
	var library model.Library
	err := l.dbConnection.Preload("Books").First(&library, "library_id = ?", libraryID).Error
	if err != nil {
		return nil, err
	}
	return library.Books, nil
}

func NewLibraryRepository(db *gorm.DB) ILibraryRepository {
	return &libraryRepository{
		dbConnection: db,
	}
}

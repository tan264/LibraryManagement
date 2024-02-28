package service

import (
	"LibraryManagement/model"
	"LibraryManagement/repository"
)

type ILibraryService interface {
	GetBooks(libraryID uint) ([]model.Book, error)
}

type libraryService struct {
	libraryRepository repository.ILibraryRepository
}

func (l libraryService) GetBooks(libraryID uint) ([]model.Book, error) {
	books, err := l.libraryRepository.GetBooks(libraryID)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func NewLibraryService(libraryRepository repository.ILibraryRepository) ILibraryService {
	return libraryService{
		libraryRepository: libraryRepository,
	}
}

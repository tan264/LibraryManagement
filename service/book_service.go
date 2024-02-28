package service

import (
	"LibraryManagement/model"
	"LibraryManagement/repository"
	"encoding/csv"
	"github.com/jinzhu/copier"
	"log"
	"mime/multipart"
	"strconv"
)

type IBookService interface {
	CreateBook(request model.CreateBookRequest) (model.Book, error)
	UpdateBook(request model.UpdateBookRequest) (model.Book, error)
	DeleteBook(bookID uint) error
	GetBookByID(bookID uint) (model.Book, error)
	FilterBook(request model.FilterBookRequest) ([]model.Book, error)
	ImportBookData(file *multipart.FileHeader) (string, error)
}

type bookService struct {
	bookRepository repository.IBookRepository
}

func (b bookService) UpdateBook(request model.UpdateBookRequest) (model.Book, error) {
	bookToUpdate, err := b.bookRepository.GetBookByID(request.BookID)
	if err != nil {
		return model.Book{}, err
	}
	if request.Title != nil {
		bookToUpdate.Title = *request.Title
	}
	if request.LibraryID != 0 {
		bookToUpdate.LibraryID = request.LibraryID
	}
	userUpdated, err := b.bookRepository.UpdateBook(bookToUpdate)
	if err != nil {
		return model.Book{}, err
	}
	return userUpdated, nil
}

func (b bookService) DeleteBook(bookID uint) error {
	return b.bookRepository.DeleteBook(bookID)
}

func (b bookService) GetBookByID(bookID uint) (model.Book, error) {
	bookFound, err := b.bookRepository.GetBookByID(bookID)
	if err != nil {
		return model.Book{}, err
	}
	return bookFound, nil
}

func (b bookService) FilterBook(request model.FilterBookRequest) ([]model.Book, error) {
	books, err := b.bookRepository.FilterBook(request)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b bookService) CreateBook(request model.CreateBookRequest) (model.Book, error) {
	bookToCreate := model.Book{}
	err := copier.Copy(&bookToCreate, &request)
	if err != nil {
		return model.Book{}, err
	}
	createdBook, err := b.bookRepository.CreateBook(bookToCreate)
	if err != nil {
		return model.Book{}, err
	}
	return createdBook, nil
}

func (b bookService) ImportBookData(file *multipart.FileHeader) (string, error) {
	openedFile, err := file.Open()
	if err != nil {
		return "Failed to open file", err
	}
	defer func(openedFile multipart.File) {
		err = openedFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(openedFile)
	reader := csv.NewReader(openedFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return "Failed to read file", err
	}
	for _, record := range records {
		libraryID, err := strconv.ParseUint(record[1], 10, 64)
		if err != nil {
			log.Println(err)
		}
		bookToCreate := model.Book{Title: record[0], LibraryID: uint(libraryID)}
		_, err = b.bookRepository.CreateBook(bookToCreate)
		if err != nil {
			log.Println(err)
		}
	}
	return "Successfully imported", nil
}

func NewBookService(bookRepository repository.IBookRepository) IBookService {
	return bookService{
		bookRepository: bookRepository,
	}
}

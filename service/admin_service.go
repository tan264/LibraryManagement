package service

import (
	"LibraryManagement/model"
	"LibraryManagement/repository"
	"encoding/csv"
	"log"
	"mime/multipart"
	"strconv"
)

type IAdminService interface {
	ImportBookData(header *multipart.FileHeader) (string, error)
}

type adminService struct {
	bookRepository repository.IBookRepository
}

func NewAdminService(bookRepository repository.IBookRepository) IAdminService {
	return adminService{
		bookRepository: bookRepository,
	}
}

func (a adminService) ImportBookData(file *multipart.FileHeader) (string, error) {
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
		_, err = a.bookRepository.CreateBook(bookToCreate)
		if err != nil {
			log.Println(err)
		}
	}
	return "Successfully imported", nil
}

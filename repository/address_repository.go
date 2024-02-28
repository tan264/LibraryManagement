package repository

import (
	"LibraryManagement/model"
	"gorm.io/gorm"
)

type IAddressRepository interface {
	GetLibraries(name string) ([]model.Library, error)
}

type addressRepository struct {
	dbConnection *gorm.DB
}

func (a addressRepository) GetLibraries(name string) ([]model.Library, error) {
	var address model.Address
	err := a.dbConnection.Preload("Libraries").First(&address, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return address.Libraries, nil
}

func NewAddressRepository(db *gorm.DB) IAddressRepository {
	return &addressRepository{
		dbConnection: db,
	}
}

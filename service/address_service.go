package service

import (
	"LibraryManagement/model"
	"LibraryManagement/repository"
)

type IAddressService interface {
	GetLibraries(name string) ([]model.Library, error)
}

type addressService struct {
	addressRepository repository.IAddressRepository
}

func (a addressService) GetLibraries(name string) ([]model.Library, error) {
	libraries, err := a.addressRepository.GetLibraries(name)
	if err != nil {
		return nil, err
	}
	return libraries, nil
}

func NewAddressService(addressRepository repository.IAddressRepository) IAddressService {
	return addressService{
		addressRepository: addressRepository,
	}
}

package service

import (
	"LibraryManagement/model"
	"LibraryManagement/repository"
)

type IUserService interface {
	EditAccount(request model.EditAccountRequest, userID uint) (*model.User, error)
	DeleteAccount(userID uint) error
	FilterAccount(request model.FilterUserRequest) ([]model.User, error)
	StatisticizeAccountByCreatedAt() (map[string]uint, error)
	CheckoutBook(bookID uint, userID uint) (model.Checkout, error)
	ReturnBook(bookID uint, userID uint) (model.Checkout, error)
}

type UserService struct {
	userRepository     repository.IUserRepository
	checkoutRepository repository.ICheckoutRepository
}

func (u UserService) CheckoutBook(bookID uint, userID uint) (model.Checkout, error) {
	checkoutBook, err := u.checkoutRepository.CheckoutBook(bookID, userID)
	if err != nil {
		return checkoutBook, err
	}
	return checkoutBook, nil
}

func (u UserService) ReturnBook(bookID uint, userID uint) (model.Checkout, error) {
	returnedBook, err := u.checkoutRepository.ReturnBook(bookID, userID)
	if err != nil {
		return returnedBook, err
	}
	return returnedBook, nil
}

func (u UserService) FilterAccount(request model.FilterUserRequest) (users []model.User, err error) {
	users, err = u.userRepository.FilterUser(request)
	if err != nil {
		return users, err
	}
	return users, nil
}

func NewUserService(userRepository repository.IUserRepository, checkoutRepository repository.ICheckoutRepository) IUserService {
	return UserService{
		userRepository:     userRepository,
		checkoutRepository: checkoutRepository,
	}
}

func (u UserService) StatisticizeAccountByCreatedAt() (result map[string]uint, err error) {
	result, err = u.userRepository.CountUserGroupByCreatedAt()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u UserService) EditAccount(request model.EditAccountRequest, userID uint) (*model.User, error) {
	userToUpdate, err := u.userRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if request.Email != nil {
		userToUpdate.Email = request.Email
	}
	if request.LastName != nil {
		userToUpdate.LastName = request.LastName
	}
	if request.FirstName != nil {
		userToUpdate.FirstName = request.FirstName
	}
	if request.Phone != nil {
		userToUpdate.Phone = request.Phone
	}
	userUpdated, err := u.userRepository.UpdateUser(*userToUpdate)
	if err != nil {
		return nil, err
	}
	return userUpdated, nil
}

func (u UserService) DeleteAccount(userID uint) (err error) {
	err = u.userRepository.DeleteUser(userID)
	return err
}

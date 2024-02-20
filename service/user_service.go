package service

import (
	"LibraryManagement/model"
	"LibraryManagement/repository"
)

type IUserService interface {
	EditAccount(request model.EditAccountRequest, userID uint) (*model.User, error)
	DeleteAccount(userID uint) model.User
}

type UserService struct {
	userRepository repository.IUserRepository
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

func (u UserService) DeleteAccount(userID uint) model.User {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return UserService{
		userRepository: userRepository,
	}
}

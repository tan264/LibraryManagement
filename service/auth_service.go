package service

import (
	"LibraryManagement/model"
	"LibraryManagement/repository"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IAuthService interface {
	RegisterUser(request model.RegisterUserRequest) model.User
	IsDuplicateUsername(username string) bool
	VerifyCredential(username string, password string) (uint, error)
}

type authService struct {
	userRepository repository.IUserRepository
}

func (a authService) VerifyCredential(username string, password string) (uint, error) {
	user, err := a.userRepository.GetByUsername(username)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return user.UserID, nil
}

func (a authService) RegisterUser(request model.RegisterUserRequest) model.User {
	userToCreate := model.User{}
	err := copier.Copy(&userToCreate, &request)
	if err != nil {
		log.Println(err.Error())
	}
	userToCreate.Password, err = hashPassword(userToCreate.Password)
	if err != nil {
		log.Println(err.Error())
	}
	result := a.userRepository.CreateUser(userToCreate)
	return result
}

func (a authService) IsDuplicateUsername(username string) bool {
	return a.userRepository.IsDuplicateUsername(username)
}

func AuthService(userRepository repository.IUserRepository) IAuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

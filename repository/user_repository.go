package repository

import (
	"LibraryManagement/model"
	"errors"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user model.User) model.User
	UpdateUser(user model.User) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByUserID(userID uint) (*model.User, error)
	IsDuplicateUsername(username string) bool
}

type userRepository struct {
	dbConnection *gorm.DB
}

func (u userRepository) UpdateUser(user model.User) (*model.User, error) {
	err := u.dbConnection.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) GetByUserID(userID uint) (*model.User, error) {
	var user model.User
	result := u.dbConnection.Where("user_id = ?", userID).Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		dbConnection: db,
	}
}

func (u userRepository) CreateUser(user model.User) model.User {
	u.dbConnection.Save(&user)
	return user
}

func (u userRepository) IsDuplicateUsername(username string) bool {
	err := u.dbConnection.Where("username = ?", username).First(&model.User{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (u userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	result := u.dbConnection.Where("username = ?", username).Take(&user)
	if result.Error == nil {
		return &user, nil
	}
	return nil, result.Error
}

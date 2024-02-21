package repository

import (
	"LibraryManagement/model"
	"database/sql"
	"errors"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user model.User) model.User
	UpdateUser(user model.User) (*model.User, error)
	DeleteUser(userID uint) error
	GetByUsername(username string) (*model.User, error)
	GetByUserID(userID uint) (*model.User, error)
	FilterUser(userToFilter model.FilterUserRequest) ([]model.User, error)
	CountUserGroupByCreatedAt() (map[string]uint, error)
	IsDuplicateUsername(username string) bool
}

type userRepository struct {
	dbConnection *gorm.DB
}

func UserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		dbConnection: db,
	}
}

func (u userRepository) CountUserGroupByCreatedAt() (result map[string]uint, err error) {
	result = make(map[string]uint)
	rows, err := u.dbConnection.Table("users").Select("date(created_at) as date, count(*) as count").Group("date(created_at)").Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var dateString string
		var amount uint
		err := rows.Scan(&dateString, &amount)
		if err != nil {
			return nil, err
		}
		result[dateString[0:10]] = amount
	}

	return result, nil
}

func (u userRepository) FilterUser(userToFilter model.FilterUserRequest) ([]model.User, error) {
	var users []model.User
	query := u.dbConnection.Find(&users)
	if userToFilter.UserID > 0 {
		query = query.Where("user_id = ?", userToFilter.UserID)
	}
	if userToFilter.Phone != "" {
		query = query.Where("phone = ?", userToFilter.Phone)
	}
	if userToFilter.CreatedAt != "" {
		query = query.Where("DATE(created_at) = ?", userToFilter.CreatedAt)
	}
	err := query.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u userRepository) DeleteUser(userID uint) (err error) {
	err = u.dbConnection.Delete(&model.User{}, userID).Error
	return err
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

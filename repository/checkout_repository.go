package repository

import (
	"LibraryManagement/model"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type ICheckoutRepository interface {
	CheckoutBook(bookID uint, userID uint) (model.Checkout, error)
	ReturnBook(bookID uint, userID uint) (model.Checkout, error)
	GetCheckoutBookByUserIDAndBookID(bookID uint, userID uint) (model.Checkout, error)
	GetAllCheckout() ([]model.Checkout, error)
	CountCheckoutByCurrentYearGroupByMonth() (map[string][2]uint, error)
}

type checkoutRepository struct {
	dbConnection *gorm.DB
}

func (c checkoutRepository) CountCheckoutByCurrentYearGroupByMonth() (map[string][2]uint, error) {
	result := make(map[string][2]uint)
	rows, err := c.dbConnection.Table("checkouts").
		Select("month(start_time) as month, "+
			"count(case when is_returned = 0 then 1 end) as borrow_amount, "+
			"count(case when is_returned = 1 then 1 end) as return_amount").
		Where("year(start_time) = ?", time.Now().Year()).
		Group("month(start_time)").
		Order("month(start_time) asc").Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var monthString string
		var borrowAmount uint
		var returnAmount uint
		err := rows.Scan(&monthString, &borrowAmount, &returnAmount)
		if err != nil {
			return nil, err
		}
		result[monthString] = [2]uint{borrowAmount, returnAmount}
	}
	return result, nil
}

func (c checkoutRepository) GetAllCheckout() ([]model.Checkout, error) {
	var checkouts []model.Checkout
	currentYear := time.Now().Year()
	err := c.dbConnection.Find(&checkouts, "year(start_time) = ?", currentYear).Error
	if err != nil {
		return []model.Checkout{}, err
	}
	return checkouts, nil
}

func (c checkoutRepository) GetCheckoutBookByUserIDAndBookID(bookID uint, userID uint) (model.Checkout, error) {
	var checkout model.Checkout
	err := c.dbConnection.First(&checkout, "book_id = ? AND user_id = ?", bookID, userID).Error
	if err != nil {
		return model.Checkout{}, err
	}
	return checkout, nil
}

func (c checkoutRepository) CheckoutBook(bookID uint, userID uint) (model.Checkout, error) {
	startTime := time.Now().Add(time.Hour * 7)
	endTime := startTime.AddDate(0, 0, 7)
	checkout := model.Checkout{
		BookID:     bookID,
		UserID:     userID,
		StartTime:  startTime,
		EndTime:    endTime,
		IsReturned: 0,
	}
	err := c.dbConnection.Save(&checkout).Error
	if err != nil {
		return model.Checkout{}, err
	}
	return checkout, nil
}

func (c checkoutRepository) ReturnBook(bookID uint, userID uint) (model.Checkout, error) {
	checkoutBook, err := c.GetCheckoutBookByUserIDAndBookID(bookID, userID)
	if err != nil {
		return model.Checkout{}, err
	}
	checkoutBook.IsReturned = 1
	err = c.dbConnection.Save(&checkoutBook).Error
	if err != nil {
		return model.Checkout{}, err
	}
	return checkoutBook, nil
}

func NewCheckoutRepository(db *gorm.DB) ICheckoutRepository {
	return checkoutRepository{
		dbConnection: db,
	}
}

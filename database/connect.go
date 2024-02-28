package database

import (
	"LibraryManagement/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() *gorm.DB {
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, pass, host, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
		return nil
	}

	log.Println("Connected to MySQL:", db.Name())
	err = db.AutoMigrate(&model.User{}, &model.Address{}, &model.Library{}, &model.Book{}, &model.Checkout{}, &model.Token{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}

package models

import (
	"github.com/fanatic75/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetBookById(Id int64) (Book, *gorm.DB) {
	var book Book
	d := db.Where("id=?", Id).Find(&book)
	return book, d
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func CreateBook(book *Book) {
	db.Create(book)
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Find("id=?", Id).Delete(&book)
	return book
}

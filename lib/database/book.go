package database

import (
	"rest/api/config"
	"rest/api/models"
)

////////////////BOOK///////////////////////
func GetBooks() (interface{}, error) {
	var book []models.Book
	if err := config.DB.Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetOneBook(id int) (models.Book, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id=?", id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func CreateBook(add_book models.Book) (interface{}, error) {

	if err := config.DB.Save(&add_book).Error; err != nil {
		return nil, err
	}
	return add_book, nil
}

func EditBook(bookSave models.Book) (interface{}, error) {
	if err := config.DB.Save(&bookSave).Error; err != nil {
		return nil, err
	}
	return bookSave, nil
}

func DeleteBook(id int) (interface{}, error) {
	var delete_book models.Book
	if err := config.DB.Find(&delete_book, "id=?", id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&delete_book).Error; err != nil {
		return nil, err
	}
	return delete_book, nil
}

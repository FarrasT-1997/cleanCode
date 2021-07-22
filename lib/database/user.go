package database

import (
	"rest/api/config"
	"rest/api/middlewares"
	"rest/api/models"
)

////////////////USER///////////////////////
func GetUsers() (interface{}, error) {
	var user []models.User
	if err := config.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(add_user models.User) (interface{}, error) {

	if err := config.DB.Save(&add_user).Error; err != nil {
		return nil, err
	}
	return add_user, nil
}

func Edituser(userSave models.User) (interface{}, error) {
	if err := config.DB.Save(&userSave).Error; err != nil {
		return nil, err
	}
	return userSave, nil
}

func DeleteUser(id int) (interface{}, error) {
	var delete_user models.User
	if err := config.DB.Find(&delete_user, "id=?", id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&delete_user).Error; err != nil {
		return nil, err
	}
	return delete_user, nil
}

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

/////////////////////JWT/////////////////////
func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User
	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(email, password string) (interface{}, error) {
	var user models.User
	var err error
	if err = config.DB.Where("email=? AND password=?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}

package database

import (
	"rest/api/config"
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

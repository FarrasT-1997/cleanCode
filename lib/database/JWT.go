package database

import (
	"rest/api/config"
	"rest/api/middlewares"
	"rest/api/models"
)

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

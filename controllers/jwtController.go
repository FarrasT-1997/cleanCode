package controllers

import (
	"net/http"
	"rest/api/lib/database"
	"rest/api/middlewares"
	"rest/api/models"
	"strconv"

	"github.com/labstack/echo"
)

////////////////JWT//////////////////////////////
func LoginUserController(c echo.Context) error {
	userData := models.User{}
	c.Bind(&userData)

	users, e := database.LoginUsers(userData.Email, userData.Password)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	users, err := database.GetDetailUsers(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"users":   users,
	})
}

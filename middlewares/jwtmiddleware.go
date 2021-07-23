package middlewares

import (
	"fmt"
	"rest/api/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//Make New Token when user login
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = int(userId)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

//Search userId based on token for authorization
func ExtractTokenUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	fmt.Println(user.Valid)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		return userId
	}
	return 0
}

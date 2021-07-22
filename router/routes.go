package router

import (
	"rest/api/constant"
	"rest/api/controllers"

	// "rest/api/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	e.GET("/users/", controllers.GetUserControllers)
	e.GET("/users/:id", controllers.GetOneUserControllers)
	e.POST("/users/", controllers.CreateUserControllers)
	e.PUT("/users/:id", controllers.EditUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)

	e.GET("/books/", controllers.GetBooksControllers)
	e.GET("/books/:id", controllers.GetOneBookControllers)
	e.POST("/books/", controllers.CreateBookControllers)
	e.PUT("/books/:id", controllers.EditBooksControllers)
	e.DELETE("/books/:id", controllers.DeleteBookControllers)

	e.POST("/login", controllers.LoginUserController)
	//jwt auth
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("/users/:id", controllers.GetUserDetailControllers)
	eJwt.GET("/users/", controllers.GetUserControllers)
	eJwt.DELETE("/users/:id", controllers.DeleteUserControllers)
	eJwt.PUT("/users/:id", controllers.EditUserControllers)

	// basic auth
	// eAuth := e.Group("")
	// eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDb))
	// eAuth.GET("/users/:id", controllers.GetOneUserControllers)
}

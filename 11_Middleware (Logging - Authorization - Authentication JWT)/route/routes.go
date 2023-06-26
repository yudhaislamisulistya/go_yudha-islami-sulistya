package route

import (
	c "project/constants"
	"project/controller"
	m "project/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(c.SECRET_JWT)))

	// user routing
	eJwt.GET("/users", controller.GetUsersController)
	eJwt.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)

	// book routing
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)
	m.LogMiddleware(e)

	// add login to get token
	e.POST("/login", controller.LoginUserController)

	return e
}

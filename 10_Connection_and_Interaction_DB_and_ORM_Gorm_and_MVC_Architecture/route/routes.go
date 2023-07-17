package route

import (
	"project/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// user routing
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)

	// book routing
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.PUT("/books/:id", controller.UpdateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)

	return e
}

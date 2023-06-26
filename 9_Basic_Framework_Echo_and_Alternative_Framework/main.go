package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// get all user
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all user",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id := c.Param("id")

	for _, item := range users {
		if strconv.Itoa(item.Id) == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user by id",
				"user":     item,
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"messages": "user not found",
	})
}

// delete user
func DeleteUserController(c echo.Context) error {
	id := c.Param("id")

	for index, item := range users {
		if strconv.Itoa(item.Id) == id {
			users = append(users[:index], users[index+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user by id",
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"messages": "user not found",
	})
}

// update user
func UpdateUserController(c echo.Context) error {
	id := c.Param("id")

	for index, item := range users {
		if strconv.Itoa(item.Id) == id {
			users[index].Name = c.FormValue("name")
			users[index].Email = c.FormValue("email")
			users[index].Password = c.FormValue("password")

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user by id",
				"user":     users[index],
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"messages": "user not found",
	})
}

// create user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.Logger.Fatal(e.Start(":8000"))

}

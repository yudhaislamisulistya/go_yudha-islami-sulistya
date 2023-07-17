package controller

import (
	"net/http"

	"project/config"
	"project/model"

	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []model.Book

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)
	config.DB.Save(&book)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id := c.Param("id")
	var book model.Book

	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    book,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	var book model.Book

	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	c.Bind(&book)

	config.DB.Save(&book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book by id",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	var book model.Book

	result := config.DB.Where("id = ?", id).Delete(&book)

	if result.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id",
	})
}

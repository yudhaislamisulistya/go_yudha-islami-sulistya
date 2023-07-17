package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/model"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetBooksController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions for success get all books
	if assert.NoError(t, GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	} else {
		assert.Fail(t, "Unexpected error occurred while getting all books")
	}
}

func TestCreateBookController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	book := model.Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Test Author",
	}
	bookJSON, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(bookJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateBookController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetBookController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	// Assertions for error get book by id
	req = httptest.NewRequest(http.MethodGet, "/books/100", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := GetBookController(c)
	// Assertions for error get book
	if assert.Error(t, GetBookController(c)) {
		e, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, e.Code)
		assert.Equal(t, "record not found", e.Message)
	}

}

func TestUpdateBookController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	book := model.Book{
		Title:  "Test Book",
		Author: "Test Author",
	}
	bookJSON, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(string(bookJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	// Assertions for error update book by id
	req = httptest.NewRequest(http.MethodPut, "/books/100", strings.NewReader(string(bookJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := UpdateBookController(c)
	// Assertions for error get book
	if assert.Error(t, UpdateBookController(c)) {
		e, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, e.Code)
		assert.Equal(t, "record not found", e.Message)
	}
}

func TestDeleteBookController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	// Assertions for error delete book by id
	req = httptest.NewRequest(http.MethodDelete, "/books/100", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := DeleteBookController(c)
	// Assertions for error get book
	if assert.Error(t, DeleteBookController(c)) {
		e, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, e.Code)
	}
}

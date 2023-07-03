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

func TestGetUsersController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateUserController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	user := model.User{
		Name:     "Test User",
		Email:    "Test@email.com",
		Password: "123456",
	}
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetUserController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/users/100", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := GetUserController(c)

	// Assertions
	if assert.Error(t, GetBookController(c)) {
		e, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, e.Code)
		assert.Equal(t, "record not found", e.Message)
	}
}

func TestUpdateUserController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	user := model.User{
		Name:     "Test User Updated",
		Email:    "Test@email.com",
		Password: "123456",
	}
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(string(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req = httptest.NewRequest(http.MethodPut, "/users/100", strings.NewReader(string(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := UpdateUserController(c)

	// Assertions
	if assert.Error(t, UpdateUserController(c)) {
		e, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, e.Code)
		assert.Equal(t, "record not found", e.Message)
	}
}

func TestDeleteUserController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req = httptest.NewRequest(http.MethodDelete, "/users/100", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := DeleteUserController(c)

	// Assertions
	if assert.Error(t, DeleteUserController(c)) {
		e, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, e.Code)
	}
}

func TestLoginUserController(t *testing.T) {
	// Setup
	config.InitDB()
	e := echo.New()
	user := model.User{
		Name:     "Test User",
		Email:    "Test@email.com",
		Password: "123456",
	}
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	} else {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

}

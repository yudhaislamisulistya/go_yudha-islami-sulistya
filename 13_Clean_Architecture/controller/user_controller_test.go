package controller

import (
	"belajar-go-echo/model"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// MockDB is a mock implementation of gorm.DB
type MockDB struct {
	mock.Mock
}

type MockContext struct {
	mock.Mock
	echo.Context
}

type MockTokenGenerator struct {
	mock.Mock
}

func (m *MockContext) Request() *http.Request {
	args := m.Called()
	return args.Get(0).(*http.Request)
}

func (m *MockContext) Response() *httptest.ResponseRecorder {
	args := m.Called()
	return args.Get(0).(*httptest.ResponseRecorder)
}

func (m *MockContext) Param(name string) string {
	args := m.Called(name)
	return args.String(0)
}

func (m *MockContext) Bind(i interface{}) error {
	args := m.Called(i)
	return args.Error(0)
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	mockArgs := m.Called(query, args)
	return mockArgs.Get(0).(*gorm.DB)
}

func (m *MockTokenGenerator) CreateToken(userId int, name string) (string, error) {
	args := m.Called(userId, name)
	return args.String(0), args.Error(1)
}

func TestGetAllUsers(t *testing.T) {
	// Create a new instance of Echo framework
	e := echo.New()

	// Create a mock database
	mockDB := new(MockDB)

	// Define the expected behavior of the mock database
	expectedUsers := []model.User{}
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{})

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler function
	handler := GetAllUsers(mockDB)
	err := handler(c)

	// Assertions

	// if nil, it means no error

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := echo.Map{
		"data": expectedUsers,
	}
	jsonString, err := toJSONString(expectedResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, rec.Body.String(), jsonString)

	// Assert that the mock database method was called
	mockDB.AssertCalled(t, "Find", mock.Anything, mock.Anything)

	// Create a new instance of Echo framework
	e = echo.New()

	// Create a mock database
	mockDB = new(MockDB)

	// Define the expected behavior of the mock database
	expectedError := errors.New("database error")
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{
		Error: expectedError,
	})

	requestData := `{
			"email": "yudha@gmail.com",
		}`

	// Create a new HTTP request
	req = httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(requestData))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	// Call the handler function
	handler = GetAllUsers(mockDB)
	err = handler(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	expectedResponse = echo.Map{
		"error": expectedError.Error(),
	}
	jsonString, _ = toJSONString(expectedResponse)
	assert.JSONEq(t, rec.Body.String(), jsonString)

	// Assert that the mock database method was called
	mockDB.AssertCalled(t, "Find", mock.Anything, mock.Anything)
}

func TestCreateUser(t *testing.T) {
	// Create a new instance of Echo framework
	e := echo.New()

	// Create a mock database
	mockDB := new(MockDB)

	// Define the expected behavior of the mock database
	expectedUser := model.User{
		Name:     "",
		Email:    "",
		Password: "",
	}

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})
	mockDB.On("Where", mock.Anything, mock.Anything).Return(&gorm.DB{})
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{})
	mockDB.On("Error").Return(nil)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler function
	handler := CreateUser(mockDB)
	err := handler(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	expectedResponse := echo.Map{
		"data": expectedUser,
	}
	jsonString, _ := toJSONString(expectedResponse)
	assert.JSONEq(t, rec.Body.String(), jsonString)

	// Assert that the mock database method was called
	mockDB.AssertCalled(t, "Create", mock.Anything, mock.Anything)
}

func TestBadRequestCreateUser(t *testing.T) {
	// Create a new instance of Echo framework
	e := echo.New()

	// Create a mock database
	mockDB := new(MockDB)

	// Define the expected behavior of the mock database
	expectedError := errors.New("database error")
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{
		Error: expectedError,
	})

	requestData := `{
		"email": "
	}`

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(requestData))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler function
	handler := CreateUser(mockDB)
	err := handler(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	expectedResponse := echo.Map{
		"error": expectedError.Error(),
	}
	jsonString, _ := toJSONString(expectedResponse)
	assert.JSONEq(t, rec.Body.String(), jsonString)

	// Assert that the mock database method was called
	mockDB.AssertCalled(t, "Create", mock.Anything, mock.Anything)
}

func toJSONString(v interface{}) (string, error) {
	// Mengonversi objek menjadi JSON string
	jsonString, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

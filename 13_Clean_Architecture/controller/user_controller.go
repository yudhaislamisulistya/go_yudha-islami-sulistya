package controller

import (
	"belajar-go-echo/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}

func GetAllUsers(db DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := make([]model.User, 0)
		err := db.Find(&users).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": users,
		})
	}
}

func CreateUser(db DB) echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		err := db.Create(&user).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusCreated, echo.Map{
			"data": user,
		})
	}
}

// func LoginUserController(db DB) echo.HandlerFunc {
// 	user := model.User{}
// 	return func(c echo.Context) error {
// 		if err := c.Bind(&user); err != nil {
// 			return c.JSON(http.StatusInternalServerError, echo.Map{
// 				"error": err.Error(),
// 			})
// 		}
// 		err := db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, echo.Map{
// 				"error": err.Error(),
// 			})
// 		}
// 		token, err := middleware.CreateToken(user.ID, user.Name)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, echo.Map{
// 				"error": err.Error(),
// 			})
// 		}
// 		userResponse := model.UserResponse{
// 			ID:    user.ID,
// 			Name:  user.Name,
// 			Email: user.Email,
// 			Token: token,
// 		}
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"data": userResponse,
// 		})
// 	}
// }

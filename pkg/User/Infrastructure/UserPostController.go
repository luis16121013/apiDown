package Infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/luis16121013/apiDown/pkg/User/Application"
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

type UserPostController struct{}

func NewUserPostController() *UserPostController {
	return &UserPostController{}
}

func (uc *UserPostController) GetUsers(s Domain.UserService) func(c echo.Context) error {
	return func(c echo.Context) error {
		users, err := Application.FindAllUser(s)
		if err != nil {
			return c.HTML(200, "<h1>NO FOUND USER!</h1>")
		}
		return c.JSON(200, users)
	}
}

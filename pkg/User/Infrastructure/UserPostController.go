package Infrastructure

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/luis16121013/apiDown/pkg/Response"
	"github.com/luis16121013/apiDown/pkg/User/Application"
	"github.com/luis16121013/apiDown/pkg/User/Domain"
)

type UserPostController struct{}

func NewUserPostController() *UserPostController {
	return &UserPostController{}
}

func (uc *UserPostController) GetUsers(s Domain.UserService) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := Response.NewResponseData()

		users, err := Application.FindAllUser(s)
		if err != nil {
			return c.HTML(http.StatusOK, "")
		}
		resp.Data = users
		return c.JSON(http.StatusOK, resp)
	}
}

func (uc *UserPostController) Login(s Domain.UserService) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := Response.NewResponseData()
		u := new(Domain.User)
		if err := c.Bind(u); err != nil{
			return c.JSON(http.StatusUnprocessableEntity,"UNPROCESSABLE ENTITY!")
		}
		
		user, err := Application.ValidateLogin(s,u)
		if err != nil {
			fmt.Println(err)
		}
		if user.Id == 0{
			resp.ResponseUnauthorized()
			return c.JSON(http.StatusUnauthorized, resp)
		}
		resp.Data = user
		return c.JSON(http.StatusOK, resp)
	}
}

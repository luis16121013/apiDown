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

// ROUTE GET-USERS
func (uc *UserPostController) GetUsers(s Domain.UserService) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := Response.NewResponseData()

		users, err := Application.FindAllUser(s)
		if err != nil {
			return c.HTML(http.StatusOK, "")
		}
		resp.Data = users
		return resp.SendData(c)
	}
}

//ROUTE GET-USER
func (uc *UserPostController) GetUser(s Domain.UserService) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := Response.NewResponseData()
		user, err := Application.FindIdUser(s, c.Param("id"))
		if err != nil {
			fmt.Println(err)
		}

		if user.Id == 0 {
			return resp.ResponseNotFound(c)
		}
		resp.Data = user
		return resp.SendData(c)
	}
}

//ROUTE LOGIN
//metodo encargado de controlar y responder al ingreso de usuarios
//valida errores de envio de datos
func (uc *UserPostController) Login(s Domain.UserService) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := Response.NewResponseData()
		u := new(Domain.User)
		if err := c.Bind(u); err != nil {
			return resp.ResponseUnprocessableEntity(c)
		}
		if u.Username == "" || u.Password == "" {
			return resp.ResponseUnauthorized(c)
		}

		user, err := Application.ValidateLogin(s, u)
		if err != nil {
			fmt.Println(err)
		}
		if user.Id == 0 {
			return resp.ResponseUnauthorized(c)
		}
		resp.Data = user
		return resp.SendData(c)
	}
}

//ROUTE CREATE-USERS
func (uc *UserPostController) CreateUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := Response.NewResponseData()
		u := new(Domain.User)
		if err := c.Bind(u); err != nil {
			return resp.ResponseUnprocessableEntity(c)
		}
		resp.Data = u
		return resp.SendData(c)
	}
}

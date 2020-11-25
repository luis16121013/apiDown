package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/luis16121013/apiDown/pkg/User/Domain"
	"github.com/luis16121013/apiDown/pkg/User/Infrastructure"
)

func StartServer(port string) {
	e := echo.New()

	r, err := Infrastructure.NewRepository()
	if err != nil {
		fmt.Println(err)
	}

	c := Infrastructure.NewUserPostController()
	s := Domain.NewService(r)

	e.GET("api/v1/users", c.GetUsers(s))

	e.Start(port)
}

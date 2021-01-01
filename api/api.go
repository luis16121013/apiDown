package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luis16121013/apiDown/pkg/User/Domain"
	"github.com/luis16121013/apiDown/pkg/User/Infrastructure"
)

func StartServer(port string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	r, err := Infrastructure.NewRepository()
	if err != nil {
		fmt.Println(err)
	}

	c := Infrastructure.NewUserPostController()
	s := Domain.NewService(r)

	e.Static("/", "cmd/public/")
	e.GET("api/v1/users", c.GetUsers(s))
	e.GET("api/v1/users/:id", c.GetUser(s))
	e.POST("api/v1/login", c.Login(s))
	e.POST("api/v1/users", c.CreateUser())

	e.Start(port)
}

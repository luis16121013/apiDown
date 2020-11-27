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
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PUT},
	}))

	r, err := Infrastructure.NewRepository()
	if err != nil {
		fmt.Println(err)
	}

	c := Infrastructure.NewUserPostController()
	s := Domain.NewService(r)

	e.Static("/","cmd/public/")
	e.GET("api/v1/users", c.GetUsers(s))
	e.POST("api/v1/login", c.Login(s))

	e.Start(port)
}

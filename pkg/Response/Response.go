package Response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//Create new responseData
func NewResponseData() *ResponseData {
	return &ResponseData{Status: http.StatusOK, Message: "success"}
}

//Response OK send DATA
func (r *ResponseData) SendData(c echo.Context) error {
	r.Status = http.StatusOK
	r.Message = "ok"
	return c.JSON(r.Status, r)
}

//Response NotFound
func (r *ResponseData) ResponseNotFound(c echo.Context) error {
	r.Status = http.StatusNotFound
	r.Message = "User Not Found!"
	return c.JSON(r.Status, r)
}

//Response unauthorized
func (r *ResponseData) ResponseUnauthorized(c echo.Context) error {
	r.Status = http.StatusUnauthorized
	r.Message = "usuario o contraseña incorrecto!"
	return c.JSON(http.StatusOK, r) //modifi
}

//Response UnprocessableEntity
func (r *ResponseData) ResponseUnprocessableEntity(c echo.Context) error {
	r.Status = http.StatusUnprocessableEntity
	r.Message = "Unprocessable Entity"
	return c.JSON(r.Status, r)
}

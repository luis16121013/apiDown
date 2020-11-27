package Response

import "net/http"

type ResponseData struct{
	Type int `json:"type"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func NewResponseData() *ResponseData{
	return &ResponseData{Type:http.StatusOK, Message:"success"}
}

func(r *ResponseData) ResponseUnauthorized(){
	r.Type = http.StatusUnauthorized
	r.Message = "usuario o contraseña incorrecto!"
}

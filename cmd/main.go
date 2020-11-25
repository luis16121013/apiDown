package main

import (
	"github.com/luis16121013/apiDown/api"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port != "" {
		port = ":" + port
	} else {
		port = ":3000"
	}

	api.StartServer(port)
}

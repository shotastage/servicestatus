package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func port() string {

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main screen")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

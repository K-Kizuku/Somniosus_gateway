package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func add(x, y int) int {
	return x + y
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

}

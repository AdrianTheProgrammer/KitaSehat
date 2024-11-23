package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// db, env := configs.DBConnect()

	e.GET("/hello", hello)

	e.Start(":8000")
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

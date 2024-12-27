package main

import (
	"KitaSehat_Backend/internal/factory"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	factory.InitFactory(e)

	e.Start(":8000")
}

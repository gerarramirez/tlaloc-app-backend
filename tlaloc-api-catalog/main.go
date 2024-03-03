package main

import (
	handlers "tlaloc-catalog/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	e.Logger.Fatal(e.Start(":1323"))
}

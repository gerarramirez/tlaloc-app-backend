package main

import (
	handlers "tlaloc-catalog/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=192.168.56.3 user=postgres password=admin123456 dbname=tlaloc_finance port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("error perro!!")
	}

	e := echo.New()
	e.GET("/", handlers.Home)
	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		println("error to init database")
	}

	e := echo.New()

}

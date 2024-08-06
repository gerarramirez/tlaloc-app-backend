package main

import (
	"tlaloc-catalog/dal"
	"tlaloc-catalog/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=192.168.56.3 user=postgres password=admin123456 dbname=tlaloc_finance port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("error perro!!")
	}

	e := echo.New()
	b := dal.NewBankDal(db)
	b2 := dal.NewBanksProducts(db)
	h := handler.NewHandler(b, b2)
	e.POST("/bank/create", h.Create)
	e.GET("/bank/findAll", h.FindAll)
	e.POST("/bank/update", h.Update)
	e.POST("/bank-products/create", h.CreateBanksProduct)
	e.GET("/bank-products/findAll", h.FindAllBanksProducts)
	e.POST("/bank-products/update", h.UpdateBanksProducts)
	e.Logger.Fatal(e.Start(":1323"))
}

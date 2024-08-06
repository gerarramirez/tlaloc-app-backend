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
	cc := dal.NewCommercesCategories(db)
	cs := dal.NewCommercesSubcategories(db)
	c := dal.NewCommercesDal(db)
	ec := dal.NewExpensesCategories(db)
	h := handler.NewHandler(b, b2, cc, cs, c, ec)

	e.POST("/expenses-categories/create", h.CreateExpensesCategories)
	e.POST("/expenses-categories/update", h.UpdateExpensesCategories)
	e.GET("/expenses-categories/findAll", h.FindAllExpensesCategories)

	e.POST("/commerce-categories/create", h.CreateCommercesCategories)
	e.POST("/commerce-categories/update", h.UpdateCommercesCategories)
	e.GET("/commerce-categories/findAll", h.FindAllCommercesCategories)

	e.POST("/commerce-subcategories/create", h.CreateCommercesSubcategories)
	e.POST("/commerce-subcategories/update", h.UpdateCommercesSubcategories)
	e.GET("/commerce-subcategories/findAll", h.FindAllCommercesSubcategories)

	e.POST("/commerces/create", h.CreateCommerce)
	e.POST("/commerces/update", h.UpdateCommerces)
	e.GET("/commerces/findAll", h.FindAllCommerces)

	e.POST("/bank/create", h.Create)
	e.GET("/bank/findAll", h.FindAll)
	e.POST("/bank/update", h.Update)
	e.POST("/bank-products/create", h.CreateBanksProduct)
	e.GET("/bank-products/findAll", h.FindAllBanksProducts)
	e.POST("/bank-products/update", h.UpdateBanksProducts)
	e.Logger.Fatal(e.Start(":1323"))
}

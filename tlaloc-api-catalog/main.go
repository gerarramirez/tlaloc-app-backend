package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"tlaloc-catalog/dal"
	"tlaloc-catalog/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
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
	exp := dal.NewExpensesDal(db)
	it := dal.NewDalIncomeType(db)
	pt := dal.NewProductTypesDAO(db)
	h := handler.NewHandler(b, b2, cc, cs, c, ec, exp, it, pt)

	e.POST("/expenses/create", h.CreateExpenses)
	e.POST("/expenses/update", h.UpdateExpenses)
	e.GET("/expenses/findAll", h.FindAllExpenses)

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

	e.POST("/income-types/create", h.CreateIncomeTypes)
	e.POST("/income-types/update", h.UpdateIncomeTypes)
	e.GET("/income-types/findAll", h.FindAllIncomeType)

	e.POST("/product-types/create", h.CreateProductType)
	e.GET("/product-types/GetAll", h.GetAllProductType)
	e.POST("/product-types/update", h.UpdateProductType)

	e.Logger.Fatal(e.Start(":1323"))
}

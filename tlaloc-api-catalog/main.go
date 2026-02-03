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
	ir := dal.NewInterestRate(db)
	h := handler.NewHandler(b, b2, cc, cs, c, ec, exp, it, pt, ir)
	RegisterRoutes(e, h)
	e.Logger.Fatal(e.Start(":1323"))
}

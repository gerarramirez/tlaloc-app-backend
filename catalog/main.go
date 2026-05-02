package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"tlaloc-catalog-service/dal"
	"tlaloc-catalog-service/handler"
	tokencheck "tlaloc-catalog-service/pkg/tokencheck"

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
		log.Fatalf("Failed to connect to database: %v", err)
	}
	e := echo.New()

	// Add authentication middleware
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	e.Use(tokencheck.RequireAuth(jwtSecret))

	bankDal := dal.NewBankDal(db)
	bankProductsDal := dal.NewBanksProducts(db)
	commerceCategoriesDal := dal.NewCommercesCategories(db)
	commerceSubcategoriesDal := dal.NewCommercesSubcategories(db)
	commercesDal := dal.NewCommercesDal(db)
	expensesCategoriesDal := dal.NewExpensesCategories(db)
	expensesDal := dal.NewExpensesDal(db)
	incomeTypesDal := dal.NewDalIncomeType(db)
	productTypesDal := dal.NewProductTypesDAO(db)
	interestRateDal := dal.NewInterestRate(db)
	
	catalogHandler := handler.NewHandler(
		bankDal, bankProductsDal, commerceCategoriesDal, commerceSubcategoriesDal,
		commercesDal, expensesCategoriesDal, expensesDal, incomeTypesDal,
		productTypesDal, interestRateDal,
	)
	
	RegisterRoutes(e, catalogHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

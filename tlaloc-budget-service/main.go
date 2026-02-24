package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"tlaloc-budget-service/dal"
	"tlaloc-budget-service/handler"
	tokencheck "tlaloc-budget-service/pkg/tokencheck"
)

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatal("Error loading configuration file" + err.Error())
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		println("Error en la connection")
	}

	e := echo.New()

	// Add authentication middleware
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	e.Use(tokencheck.RequireAuth(jwtSecret))

	b := dal.NewBudgetDal(db)
	budgetHandler := handler.NewHandler(b)

	e.POST("/management-budget/create", budgetHandler.CreateBudget)
	e.POST("/management-budget-expeses/categories/create", budgetHandler.CreateBudgetExpenseCate)
	e.GET("/management-budget/wholebudget/:id", budgetHandler.GetWholeBudget)

	e.Logger.Fatal(e.Start(":8080"))

}

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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
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

	budgetRepo := dal.NewBudgetDal(db)
	budgetHandler := handler.NewHandler(budgetRepo)

	e.POST("/budgets", budgetHandler.CreateBudget)
	e.POST("/budgets/expense-categories", budgetHandler.CreateBudgetExpenseCategory)
	e.GET("/budgets/:id", budgetHandler.GetBudget)

	e.Logger.Fatal(e.Start(":8080"))

}

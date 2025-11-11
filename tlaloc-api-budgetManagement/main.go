package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"tlaloc-api-budgetManagement/dal"
	"tlaloc-api-budgetManagement/handler"
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
	b := dal.NewBudgetDal(db)
	budgetHandler := handler.NewHandler(b)

	e.POST("/management-budger/create", budgetHandler.Create)

	e.Logger.Fatal(e.Start(":8080"))

}

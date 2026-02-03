package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"tlaloc-api-transaction/dal"
	"tlaloc-api-transaction/handler"
)

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatal("Error loading config file " + err.Error())
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in conn to database" + err.Error())
	}
	e := echo.New()

	// dd dao dao; dh daily handler
	dd := dal.NewDailyExpensesDao(db)
	dh := handler.NewHandler(dd)

	e.POST("/transaction/create", dh.CreateDailyExpenses)

	e.Logger.Fatal(e.Start(":8081"))

}

package handler

import (
	"net/http"
	model "tlaloc-catalog/model/db"

	"github.com/labstack/echo/v4"
)

func (handler *Handler) CreateExpensesCategories(e echo.Context) error {
	ep := new(model.ExpensesCategories)

	if error := e.Bind(ep); error != nil {
		println("error in Parse json")
		return e.JSON(http.StatusInternalServerError, "error in JSON parse")
	}
	if _, error := handler.expensesCategoriesDAO.Create(ep); error != nil {
		println("error in storage to Expenses cactegories")

	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

func (handler *Handler) FindAllExpensesCategories(e echo.Context) error {
	result, error := handler.expensesCategoriesDAO.FindAll()

	if error != nil {
		println("error in extract data")
		return e.JSON(http.StatusInternalServerError, "error in extract data")
	}

	return e.JSON(http.StatusOK, result)
}

func (handler *Handler) UpdateExpensesCategories(e echo.Context) error {
	ep := new(model.ExpensesCategories)

	if error := e.Bind(ep); error != nil {
		println("Error in JSON parse")
		return e.JSON(http.StatusInternalServerError, "error in JSON parse")
	}

	if _, error := handler.expensesCategoriesDAO.Update(ep); error != nil {
		return e.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

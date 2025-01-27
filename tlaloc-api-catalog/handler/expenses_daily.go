package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	model "tlaloc-catalog/model/db"
)

func (exp *Handler) CreateExpenses(c echo.Context) error {
	u := new(model.ExpensesDaily)
	if err := c.Bind(u); err != nil {
		return errors.New("Errorazo papa")
	}

	err := exp.expenses.Create(u)

	if err != nil {
		return errors.New("Error guardando Expenses")
	}

	return c.JSON(http.StatusOK, "SUCCESS")
}

func (exp *Handler) FindAllExpenses(c echo.Context) error {
	result, err := exp.expenses.FindAll()

	if err != nil {
		return errors.New("Error getting Expenses")
	}

	return c.JSON(http.StatusOK, result)
}

func (exp *Handler) UpdateExpenses(c echo.Context) error {
	expensesModel := new(model.ExpensesDaily)

	if err := c.Bind(expensesModel); err != nil {
		errors.New("Error parsing JSON from FrontEnd")
	}

	err := exp.expenses.Update(expensesModel)

	if err != nil {
		errors.New("Error Updating expenses!!")
	}

	return c.JSON(http.StatusOK, "SUCCCESS")

}

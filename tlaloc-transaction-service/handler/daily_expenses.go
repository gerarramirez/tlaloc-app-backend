package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tlaloc-transaction-service/model"
)

func (handler *Handler) CreateDailyExpenses(c echo.Context) error {

	m := new(model.DailyExpenses)

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	result := handler.dailyExpensesDao.CreateDailyExpenses(m)

	if result != nil {
		c.JSON(http.StatusInternalServerError, "Error saving budget")
	}

	return c.JSON(http.StatusOK, "SUCCESS")
}

func (handler *Handler) UpdateDailyExpenses(c echo.Context) error {
	m := new(model.DailyExpenses)

	if err := c.Bind(m); err != nil {
		c.JSON(http.StatusInternalServerError, "Error saving budget")
	}

	result := handler.dailyExpensesDao.UpdateCreateDailyExpenses(m)

	if result != nil {
		c.JSON(http.StatusInternalServerError, "Error saving budget")
	}

	return c.JSON(http.StatusOK, "SUCCRSS")

}

func (handler *Handler) getAll(c echo.Context) error {

	result, err := handler.dailyExpensesDao.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error with transaction")
	}

	return c.JSON(http.StatusOK, result)

}

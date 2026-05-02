package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tlaloc-transaction-service/model"
)

func (h *Handler) CreateDailyExpense(c echo.Context) error {
	m := new(model.DailyExpenses)

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.dailyExpensesDao.CreateDailyExpenses(m); err != nil {
		c.Logger().Errorf("Failed to create daily expense: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create daily expense"})
	}

	return c.JSON(http.StatusCreated, m)
}

func (h *Handler) UpdateDailyExpense(c echo.Context) error {
	m := new(model.DailyExpenses)

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.dailyExpensesDao.UpdateCreateDailyExpenses(m); err != nil {
		c.Logger().Errorf("Failed to update daily expense: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update daily expense"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Daily expense updated successfully"})
}

func (h *Handler) GetDailyExpenses(c echo.Context) error {
	result, err := h.dailyExpensesDao.GetAll()

	if err != nil {
		c.Logger().Errorf("Failed to get daily expenses: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve daily expenses"})
	}

	return c.JSON(http.StatusOK, result)
}

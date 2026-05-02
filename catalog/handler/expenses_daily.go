package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	model "tlaloc-catalog-service/model/db"
)

func (h *Handler) CreateExpense(c echo.Context) error {
	u := new(model.DailyExpenses)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	err := h.expenses.Create(u)

	if err != nil {
		c.Logger().Errorf("Failed to create expense: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create expense"})
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) GetExpenses(c echo.Context) error {
	result, err := h.expenses.FindAll()

	if err != nil {
		c.Logger().Errorf("Failed to get expenses: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve expenses"})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateExpense(c echo.Context) error {
	expensesModel := new(model.DailyExpenses)

	if err := c.Bind(expensesModel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	err := h.expenses.Update(expensesModel)

	if err != nil {
		c.Logger().Errorf("Failed to update expense: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update expense"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Expense updated successfully"})
}

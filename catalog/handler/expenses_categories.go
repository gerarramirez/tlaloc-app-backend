package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateExpenseCategory(c echo.Context) error {
	ep := new(model.ExpensesCategories)

	if err := c.Bind(ep); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := h.expensesCategoriesDAO.Create(ep); err != nil {
		c.Logger().Errorf("Failed to create expense category: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create expense category"})
	}

	return c.JSON(http.StatusCreated, ep)
}

func (h *Handler) GetExpenseCategories(c echo.Context) error {
	result, err := h.expensesCategoriesDAO.FindAll()

	if err != nil {
		c.Logger().Errorf("Failed to get expense categories: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve expense categories"})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateExpenseCategory(c echo.Context) error {
	ep := new(model.ExpensesCategories)

	if err := c.Bind(ep); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.expensesCategoriesDAO.Update(ep); err != nil {
		c.Logger().Errorf("Failed to update expense category: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update expense category"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Expense category updated successfully"})
}

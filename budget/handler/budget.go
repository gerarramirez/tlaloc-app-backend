package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tlaloc-budget-service/model"
)

func (h *Handler) CreateBudget(c echo.Context) error {
	u := new(model.Budget)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	
	if err := h.budgetDao.CreateBudget(u); err != nil {
		c.Logger().Errorf("Failed to create budget: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create budget"})
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) CreateBudgetExpenseCategory(c echo.Context) error {
	b := new(model.BudgetExpenseCategories)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.budgetDao.CreateBudgetExpenseCate(b); err != nil {
		c.Logger().Errorf("Failed to create budget expense category: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create budget expense category"})
	}

	return c.JSON(http.StatusCreated, b)
}

func (h *Handler) GetBudget(c echo.Context) error {
	var id string

	if err := echo.PathParamsBinder(c).String("id", &id).BindErrors(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid budget ID format"})
	}

	result, err := h.budgetDao.GetWholeBudget(&id)

	if err != nil {
		c.Logger().Errorf("Failed to get budget: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve budget"})
	}

	return c.JSON(http.StatusOK, result)
}

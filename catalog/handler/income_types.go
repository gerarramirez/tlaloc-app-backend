package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	model "tlaloc-catalog-service/model/db"
)

func (h *Handler) CreateIncomeType(c echo.Context) error {
	it := new(model.IncomeTypes)

	if err := c.Bind(it); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.incomeTypesDAO.Create(it); err != nil {
		c.Logger().Errorf("Failed to create income type: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create income type"})
	}

	return c.JSON(http.StatusCreated, it)
}

func (h *Handler) UpdateIncomeType(c echo.Context) error {
	it := new(model.IncomeTypes)

	if err := c.Bind(it); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.incomeTypesDAO.Update(it); err != nil {
		c.Logger().Errorf("Failed to update income type: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update income type"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Income type updated successfully"})
}

func (h *Handler) GetIncomeTypes(c echo.Context) error {
	result, err := h.incomeTypesDAO.FindAll()

	if err != nil {
		c.Logger().Errorf("Failed to get income types: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve income types"})
	}

	return c.JSON(http.StatusOK, result)
}

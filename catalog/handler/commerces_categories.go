package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCommerceCategory(c echo.Context) error {
	ct := new(model.CommercesCategories)

	if err := c.Bind(ct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.commercesCategoriesDAO.Create(ct); err != nil {
		c.Logger().Errorf("Failed to create commerce category: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create commerce category"})
	}

	return c.JSON(http.StatusCreated, ct)
}

func (h *Handler) GetCommerceCategories(c echo.Context) error {
	result, err := h.commercesCategoriesDAO.FindAll()

	if err != nil {
		c.Logger().Errorf("Failed to get commerce categories: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve commerce categories"})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateCommerceCategory(c echo.Context) error {
	ct := new(model.CommercesCategories)

	if err := c.Bind(ct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.commercesCategoriesDAO.Update(ct); err != nil {
		c.Logger().Errorf("Failed to update commerce category: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update commerce category"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Commerce category updated successfully"})
}

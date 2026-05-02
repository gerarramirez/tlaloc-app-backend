package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCommerceSubcategory(c echo.Context) error {
	cs := new(model.CommercesSubcategories)
	if err := c.Bind(cs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.commercesSubcategoriesDAO.Create(cs); err != nil {
		c.Logger().Errorf("Failed to create commerce subcategory: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create commerce subcategory"})
	}

	return c.JSON(http.StatusCreated, cs)
}

func (h *Handler) GetCommerceSubcategories(c echo.Context) error {
	result, err := h.commercesSubcategoriesDAO.FindAll()

	if err != nil {
		c.Logger().Errorf("Failed to get commerce subcategories: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve commerce subcategories"})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateCommerceSubcategory(c echo.Context) error {
	cs := new(model.CommercesSubcategories)
	if err := c.Bind(cs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.commercesSubcategoriesDAO.Update(cs); err != nil {
		c.Logger().Errorf("Failed to update commerce subcategory: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update commerce subcategory"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Commerce subcategory updated successfully"})
}

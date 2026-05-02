package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	model "tlaloc-catalog-service/model/db"
)

func (h *Handler) CreateProductType(c echo.Context) error {
	pt := new(model.ProductType)
	if err := c.Bind(pt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.productTypeDAO.Create(pt); err != nil {
		c.Logger().Errorf("Failed to create product type: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product type"})
	}
	
	return c.JSON(http.StatusCreated, pt)
}

func (h *Handler) GetProductTypes(c echo.Context) error {
	result, err := h.productTypeDAO.GetAll()

	if err != nil {
		c.Logger().Errorf("Failed to get product types: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve product types"})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateProductType(c echo.Context) error {
	pt := new(model.ProductType)

	if err := c.Bind(pt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.productTypeDAO.Update(pt); err != nil {
		c.Logger().Errorf("Failed to update product type: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product type"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product type updated successfully"})
}

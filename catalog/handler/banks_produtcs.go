package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateBankProduct(c echo.Context) error {
	p := new(model.BanksProducts)

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.banksProductsDAO.Create(p); err != nil {
		c.Logger().Errorf("Failed to create bank product: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create bank product"})
	}

	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) GetBankProducts(c echo.Context) error {
	b, err := h.banksProductsDAO.FindAll()

	if err != nil {
		c.Logger().Errorf("Failed to get bank products: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve bank products"})
	}

	return c.JSON(http.StatusOK, b)
}

func (h *Handler) UpdateBankProduct(c echo.Context) error {
	bp := new(model.BanksProducts)

	if err := c.Bind(bp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.banksProductsDAO.Update(bp); err != nil {
		c.Logger().Errorf("Failed to update bank product: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update bank product"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Bank product updated successfully"})
}

package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateBank(c echo.Context) error {
	u := new(model.Bank)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	err := h.bankDAO.Create(u)
	if err != nil {
		c.Logger().Errorf("Failed to create bank: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create bank"})
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) GetBanks(c echo.Context) error {
	b, err := h.bankDAO.FindAll()
	if err != nil {
		c.Logger().Errorf("Failed to get banks: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve banks"})
	}
	return c.JSON(http.StatusOK, b)
}

func (h *Handler) UpdateBank(c echo.Context) error {
	u := new(model.Bank)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := h.bankDAO.Update(u); err != nil {
		c.Logger().Errorf("Failed to update bank: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update bank"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Bank updated successfully"})
}

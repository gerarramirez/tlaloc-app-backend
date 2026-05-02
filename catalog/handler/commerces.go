package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCommerce(c echo.Context) error {
	commerce := new(model.Commerces)

	if err := c.Bind(commerce); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if _, err := h.commercesDAO.Create(commerce); err != nil {
		c.Logger().Errorf("Failed to create commerce: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create commerce"})
	}

	return c.JSON(http.StatusCreated, commerce)
}

func (h *Handler) GetCommerces(c echo.Context) error {
	commerces, err := h.commercesDAO.FindAll()

	if err != nil {
		c.Logger().Errorf("Failed to get commerces: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve commerces"})
	}

	return c.JSON(http.StatusOK, commerces)
}

func (h *Handler) UpdateCommerce(c echo.Context) error {
	commerce := new(model.Commerces)

	if err := c.Bind(commerce); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if _, err := h.commercesDAO.Update(commerce); err != nil {
		c.Logger().Errorf("Failed to update commerce: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update commerce"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Commerce updated successfully"})
}

package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tlaloc-api-budgetManagement/model"
)

func (handler *Handler) Create(c echo.Context) error {

	u := new(model.Budget)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error parsing JSON")
	}
	error := handler.budgetDao.Create(u)

	if error != nil {
		c.JSON(http.StatusInternalServerError, "Error saving budget")
	}

	return c.JSON(http.StatusOK, "Registered of buget is done")

}

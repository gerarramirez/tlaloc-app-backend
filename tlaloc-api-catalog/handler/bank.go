package handler

import (
	"net/http"
	model "tlaloc-catalog/model/db"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func (bank *Handler) Create(c echo.Context) error {
	u := new(model.Bank)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, "Errorazo papa")
	}

	a, err := bank.bankDAO.Create(u)

	if err != nil {
		return c.String(http.StatusOK, "Hello World!")
	}

	return c.JSON(http.StatusOK, "Hello World! "+a.ID)
}

package handler

import (
	"net/http"
	model "tlaloc-catalog/model/db"

	"github.com/labstack/echo/v4"
)

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

func (bank *Handler) FindAll(c echo.Context) error {
	b, err := bank.bankDAO.FindAll()
	if err != nil {
		println("error")
	}
	return c.JSON(http.StatusOK, b)
}

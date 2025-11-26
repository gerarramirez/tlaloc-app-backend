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

	err := bank.bankDAO.Create(u)

	if err != nil {
		return c.String(http.StatusOK, "Hello World!")
	}

	return c.JSON(http.StatusOK, "Hello World! ")
}

func (bank *Handler) FindAll(c echo.Context) error {
	b, err := bank.bankDAO.FindAll()
	if err != nil {
		println("error")
	}
	return c.JSON(http.StatusOK, b)
}

func (bank *Handler) Update(c echo.Context) error {
	u := new(model.Bank)
	if err := c.Bind(u); err != nil {
		println("Error parseando el JSON")
		return err
	}
	if err := bank.bankDAO.Update(u); err != nil {
		println("Erro actualizando el dato")
		return err
	}

	return c.JSON(http.StatusOK, "Registro Actualizado de forma correcta")
}

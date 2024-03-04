package handlers

import (
	"net/http"
	"tlaloc-catalog/dal"
	model "tlaloc-catalog/model/db"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func create(c echo.Context) (model.BankEntity, error) {
	b := new(model.Bank)

	return dal.Create(c, b)
}

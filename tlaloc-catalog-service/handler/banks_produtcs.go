package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db"

	"github.com/labstack/echo/v4"
)

func (handler *Handler) CreateBanksProduct(e echo.Context) error {
	p := new(model.BanksProducts)

	if err := e.Bind(p); err != nil {
		return e.JSON(http.StatusInternalServerError, "Json con mal formación")
	}

	if error := handler.banksProductsDAO.Create(p); error != nil {
		return e.JSON(http.StatusInternalServerError, "error")
	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

func (handler *Handler) FindAllBanksProducts(c echo.Context) error {
	b, error := handler.banksProductsDAO.FindAll()

	if error != nil {
		return c.JSON(http.StatusInternalServerError, "error en el JSON")
	}

	return c.JSON(http.StatusOK, b)
}

func (handler *Handler) UpdateBanksProducts(c echo.Context) error {
	bp := new(model.BanksProducts)

	if error := c.Bind(&bp); error != nil {
		println("Error parseando modelo UpdateBanksProducts")
		return c.JSON(http.StatusInternalServerError, "Error en el paeseo del JSON")
	}

	if error := handler.banksProductsDAO.Update(bp); error != nil {
		return c.JSON(http.StatusInternalServerError, "error en el servidor")

	}

	return c.JSON(http.StatusOK, "SUCCESS")
}

package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db"

	"github.com/labstack/echo/v4"
)

func (handler *Handler) CreateCommerce(e echo.Context) error {
	c := new(model.Commerces)

	if error := e.Bind(c); error != nil {
		println("error en el parseo del modelo")
		return e.JSON(http.StatusInternalServerError, "error en la creacion de Commerces")
	}

	if _, error := handler.commercesDAO.Create(c); error != nil {
		return e.JSON(http.StatusInternalServerError, "error en la creacion de Commerces")
	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

func (handler *Handler) FindAllCommerces(e echo.Context) error {
	c, error := handler.commercesDAO.FindAll()

	if error != nil {
		e.JSON(http.StatusInternalServerError, "error en la extraccion de commerces")
	}

	return e.JSON(http.StatusOK, c)
}

func (handler *Handler) UpdateCommerces(e echo.Context) error {
	c := new(model.Commerces)

	if error := e.Bind(&c); error != nil {
		println("Error en la extraccion de commerces")
		e.JSON(http.StatusInternalServerError, "error en la actualizacion de commerce")
	}

	if _, error := handler.commercesDAO.Update(c); error != nil {
		println("error actualizando commerces")
		return e.JSON(http.StatusInternalServerError, "Error actualizando commerces")
	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

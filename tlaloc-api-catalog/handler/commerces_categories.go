package handler

import (
	"net/http"
	model "tlaloc-catalog/model/db"

	"github.com/labstack/echo/v4"
)

func (handler *Handler) CreateCommercesCategories(e echo.Context) error {
	ct := new(model.CommercesCategories)

	if error := e.Bind(ct); error != nil {
		println("error en el parseo de categorias")
		return e.JSON(http.StatusInternalServerError, "error en el parseo del modelo")
	}

	if _, error := handler.commercesCategoriesDAO.Create(ct); error != nil {
		println("error en la creacion de categorias de comercio")
		return e.JSON(http.StatusInternalServerError, "error en la creacion de las categorias de comercio")
	}

	return e.JSON(http.StatusInternalServerError, "SUCCESS")
}

func (handler *Handler) FindAllCommercesCategories(e echo.Context) error {

	result, error := handler.commercesCategoriesDAO.FindAll()

	if error != nil {
		println("error en la extraccion de categorias de commercios")
		return e.JSON(http.StatusInternalServerError, "error en la extraccion de commercios")
	}

	return e.JSON(http.StatusOK, result)

}

func (handler *Handler) UpdateCommercesCategories(e echo.Context) error {
	ct := new(model.CommercesCategories)

	if error := e.Bind(ct); error != nil {
		println("error in parse JSON commerces categories")
	}

	if _, error := handler.commercesCategoriesDAO.Create(ct); error != nil {
		println("error to create commerces categories")
	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

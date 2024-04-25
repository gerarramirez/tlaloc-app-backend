package handler

import (
	"net/http"
	model "tlaloc-catalog/model/db"

	"github.com/labstack/echo/v4"
)

func (handler *Handler) CreateCommercesSubcategories(e echo.Context) error {
	c := new(model.CommercesSubcategories)
	if error := e.Bind(c); error != nil {
		println("error in parse JSON commerces subcategories")
		return e.JSON(http.StatusInternalServerError, "error in PARSE")
	}

	if error := handler.commercesSubcategoriesDAO.Create(c); error != nil {
		println("error in create subcategories commerces")
		e.JSON(http.StatusInternalServerError, "error in subcategories commerces")
	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

func (handler *Handler) FindAllCommercesSubcategories(e echo.Context) error {
	result, error := handler.commercesSubcategoriesDAO.FindAll()

	if error != nil {
		println("error parse  JSON")
		return e.JSON(http.StatusInternalServerError, "wrong request")
	}

	return e.JSON(http.StatusOK, result)
}

func (handler *Handler) UpdateCommercesSubcategories(e echo.Context) error {
	cs := new(model.CommercesSubcategories)
	if error := e.Bind(cs); error != nil {
		println("error in  JSON parse")
		return e.JSON(http.StatusInternalServerError, "Wrong parses in JSON")
	}

	if error := handler.commercesSubcategoriesDAO.Update(cs); error != nil {
		e.JSON(http.StatusInternalServerError, "Wrong in save update")
	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

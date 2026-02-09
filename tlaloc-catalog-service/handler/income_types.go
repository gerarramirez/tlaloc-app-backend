package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	model "tlaloc-catalog-service/model/db"
)

func (incomeTypes *Handler) CreateIncomeTypes(e echo.Context) error {
	it := new(model.IncomeTypes)

	if err := e.Bind(it); err != nil {
		return errors.New("Error en el servidor")
	}

	if result := incomeTypes.incomeTypesDAO.Create(it); result != nil {
		print("Error ")
		return result
	}

	return e.JSON(http.StatusOK, "SUCCESS")

}

func (incomeTypes *Handler) UpdateIncomeTypes(e echo.Context) error {
	it := new(model.IncomeTypes)

	if err := e.Bind(it); err != nil {
		errors.New("Eror en el servidor")
	}

	if err := incomeTypes.incomeTypesDAO.Update(it); err != nil {
		return e.JSON(http.StatusInternalServerError, "Ocurrio un error")
	}

	return e.JSON(http.StatusOK, "SUCCES")
}

func (incomeTypes *Handler) FindAllIncomeType(e echo.Context) error {

	result, err := incomeTypes.incomeTypesDAO.FindAll()

	if err != nil {
		return e.JSON(http.StatusInternalServerError, "ERROR")
	}

	return e.JSON(http.StatusOK, result)
}

package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	model "tlaloc-catalog/model/db"
)

func (handler *Handler) CreateProductType(e echo.Context) error {
	pt := new(model.ProductType)
	if err := e.Bind(pt); err != nil {
		return errors.New("Error en el servidor")
	}

	if result := handler.productTypeDAO.Create(pt); result != nil {
		return errors.New("Eror en el servidor")
	}
	return e.JSON(http.StatusOK, "succes")
}

func (handler *Handler) GetAllProductType(e echo.Context) error {
	pt := new(model.ProductType)
	if err := e.Bind(pt); err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	result, err := handler.productTypeDAO.GetAll()

	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, result)
}

func (handler *Handler) UpdateProductType(e echo.Context) error {
	pt := new(model.ProductType)

	if err := e.Bind(pt); err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := handler.productTypeDAO.Update(pt); err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, "SUCCESS")
}

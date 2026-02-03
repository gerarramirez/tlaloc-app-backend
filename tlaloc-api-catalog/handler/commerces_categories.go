package handler

import (
	"net/http"
	model "tlaloc-catalog/model/db"

	"github.com/labstack/echo/v4"
)

func (handler *Handler) CreateCommercesCategories(e echo.Context) error {
	ct := new(model.CommercesCategories)

	// 1. Validar el binding
	if err := e.Bind(ct); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Formato de datos inválido"})
	}

	// 2. Intentar crear en DB
	if err := handler.commercesCategoriesDAO.Create(ct); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la categoría"})
	}

	// 3. Retornar 201 Created
	return e.JSON(http.StatusCreated, ct)
}

func (handler *Handler) FindAllCommercesCategories(e echo.Context) error {
	result, err := handler.commercesCategoriesDAO.FindAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener las categorías"})
	}

	return e.JSON(http.StatusOK, result)
}

func (handler *Handler) UpdateCommercesCategories(e echo.Context) error {
	ct := new(model.CommercesCategories)

	if err := e.Bind(ct); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Error en el parseo de datos"})
	}

	// Cambiado a .Update (asegúrate que tu DAO tenga este método)
	if err := handler.commercesCategoriesDAO.Update(ct); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar la categoría"})
	}

	return e.JSON(http.StatusOK, map[string]string{"message": "Categoría actualizada correctamente"})
}

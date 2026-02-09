package handler

import (
	"net/http"
	model "tlaloc-catalog-service/model/db" // Ajusta esta ruta a tu estructura real

	"github.com/labstack/echo/v4"
)

// CreateInterestRate maneja la creación de una nueva tasa de interés
func (handler *Handler) CreateInterestRate(e echo.Context) error {
	ir := new(model.InterestRate)

	// 1. Validar el binding del JSON
	if err := e.Bind(ir); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Formato de tasa de interés inválido"})
	}

	// 2. Persistir en la base de datos
	if err := handler.interestRateDao.Create(ir); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar la tasa de interés"})
	}

	// 3. Retornar éxito (201 Created)
	return e.JSON(http.StatusCreated, ir)
}

// FindAllInterestRates obtiene todas las tasas registradas
func (handler *Handler) FindAllInterestRates(e echo.Context) error {
	result, err := handler.interestRateDao.FindAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener las tasas de interés"})
	}

	return e.JSON(http.StatusOK, result)
}

// UpdateInterestRate actualiza una tasa existente
func (handler *Handler) UpdateInterestRate(e echo.Context) error {
	ir := new(model.InterestRate)

	if err := e.Bind(ir); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Error en el parseo de datos para actualizar"})
	}

	// Usamos el método Update del DAO
	if err := handler.interestRateDao.Update(ir); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar la tasa de interés"})
	}

	return e.JSON(http.StatusOK, map[string]string{"message": "Tasa de interés actualizada con éxito"})
}

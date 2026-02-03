package dal

import (
	"errors"
	"gorm.io/gorm"
	model "tlaloc-catalog/model/db"
)

type InterestRateDAO interface {
	Create(interestRate *model.InterestRate) error
	FindAll() ([]model.InterestRate, error)
	Update(interestRate *model.InterestRate) error
}

type InterestRate struct {
	DB *gorm.DB
}

func NewInterestRate(db *gorm.DB) *InterestRate {
	return &InterestRate{
		DB: db,
	}
}

func (dao InterestRate) Create(interestRate *model.InterestRate) error {
	// Iniciar transacción
	tx := dao.DB.Begin()

	// Usamos el puntero interestRate directamente.
	// Nota: Si el ID es autoincremental en DB, GORM lo llenará automáticamente en el struct tras el Create.
	if err := tx.Table("tlaloc_api.interest_rates").Create(interestRate).Error; err != nil {
		tx.Rollback()
		return errors.New("error en el api a la hora de guardar")
	}

	return tx.Commit().Error
}

func (dao InterestRate) FindAll() ([]model.InterestRate, error) {
	var interestRates []model.InterestRate

	// Nota: Si usas .Select("percent", "type"), el campo "id" llegará vacío al handler.
	// Es recomendable incluir el ID para poder hacer updates o deletes después.
	if err := dao.DB.Table("tlaloc_api.interest_rates").Find(&interestRates).Error; err != nil {
		return nil, errors.New("error en la extraccion de datos")
	}

	return interestRates, nil
}

func (dao InterestRate) Update(interestRate *model.InterestRate) error {
	if interestRate == nil || interestRate.Id == "" {
		return errors.New("modelo vacio o ID faltante")
	}

	tx := dao.DB.Begin()

	// .Save() requiere que el struct tenga el ID lleno para saber qué registro actualizar.
	// .Table() asegura que apunte al esquema correcto.
	if err := tx.Table("tlaloc_api.interest_rates").Save(interestRate).Error; err != nil {
		tx.Rollback()
		return errors.New("no se pudo actualizar el registro")
	}

	return tx.Commit().Error
}

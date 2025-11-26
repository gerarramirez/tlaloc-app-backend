package dal

import (
	"errors"
	"gorm.io/gorm"
	model "tlaloc-catalog/model/db"
)

type IncomeTypeDAO interface {
	Create(incomeType *model.IncomeTypes) error
	FindAll() ([]model.IncomeTypes, error)
	Update(incomeType *model.IncomeTypes) error
}

type IncomeType struct {
	DB *gorm.DB
}

func NewDalIncomeType(db *gorm.DB) *IncomeType {
	return &IncomeType{
		db,
	}
}

func (incomeTypeDb *IncomeType) Create(incomeType *model.IncomeTypes) error {

	if incomeType == nil {
		return errors.New("Datos insuficientes")
	}

	db := incomeTypeDb.DB.Begin()

	if err := db.Table("tlaloc_api.income_types").Select("Name").Create(&incomeType).Error; err != nil {
		println("PASO POR ACA")
		return errors.New("Error en el server")
	}
	return db.Commit().Error
}

func (incomeTypeDb *IncomeType) FindAll() ([]model.IncomeTypes, error) {
	var (
		result []model.IncomeTypes
	)

	if err := incomeTypeDb.DB.Begin().Table("tlaloc_api.income_types").Find(&result).Error; err != nil {
		return nil, errors.New("Error en el server")
	}

	return result, nil

}

func (incomeTypeDb *IncomeType) Update(incometype *model.IncomeTypes) error {
	if incometype == nil {
		errors.New("Datos insuficientes")
	}

	db := incomeTypeDb.DB.Begin()

	if err := db.Table("tlaloc_api.income_types").Save(incometype).Error; err != nil {
		return errors.New("Error en servidor ")
	}

	return db.Commit().Error

}

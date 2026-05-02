package dal

import (
	"errors"
	"gorm.io/gorm"
	"tlaloc-transaction-service/model"
)

type DailyExpensesDao interface {
	CreateDailyExpenses(expenses *model.DailyExpenses) error
	UpdateCreateDailyExpenses(expenses *model.DailyExpenses) error
	GetAll() ([]model.DailyExpenses, error)
}

type DailyExpensesDaoImpl struct {
	DB *gorm.DB
}

func NewDailyExpensesDao(dao *gorm.DB) *DailyExpensesDaoImpl {
	return &DailyExpensesDaoImpl{
		DB: dao,
	}
}

func (dao *DailyExpensesDaoImpl) CreateDailyExpenses(expenses *model.DailyExpenses) error {

	db := dao.DB.Begin()

	if err := db.Table("tlaloc_api.daily_expenses").Create(&expenses).Error; err != nil {
		db.Rollback()
		return errors.New("Error en el API")
	}

	return db.Commit().Error

}

func (dao *DailyExpensesDaoImpl) UpdateCreateDailyExpenses(expenses *model.DailyExpenses) error {

	db := dao.DB.Begin()

	if db.Table("tlaloc_api.daily_expenses").Save(&expenses).Error != nil {
		db.Rollback()
		return errors.New("Error al actualizar")
	}
	return db.Commit().Error
}

func (dao *DailyExpensesDaoImpl) GetAll() ([]model.DailyExpenses, error) {
	var result []model.DailyExpenses

	db := dao.DB.Begin()

	if db.Table("tlaloc_api.daily_expenses").Find(result).Error != nil {
		return nil, errors.New("")
	}

	return result, nil

}

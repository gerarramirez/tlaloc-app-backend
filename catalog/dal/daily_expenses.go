package dal

import (
	"errors"
	"gorm.io/gorm"
	model "tlaloc-catalog-service/model/db"
)

type ExpensesDao interface {
	Create(expense *model.DailyExpenses) error
	FindAll() (*[]model.DailyExpenses, error)
	Update(expenses *model.DailyExpenses) error
}

type Expenses struct {
	DB *gorm.DB
}

func NewExpensesDal(db *gorm.DB) *Expenses {
	return &Expenses{
		DB: db,
	}
}

func (expenses *Expenses) Create(exp *model.DailyExpenses) error {
	if exp == nil {
		return errors.New("modelo de expenses vacio")
	}

	model := exp

	db := expenses.DB.Begin()

	if err := db.Table("tlaloc_api.expenses").Create(model).Error; err != nil {
		db.Rollback()
		return errors.New("Error guardando expenses")
	}

	return db.Commit().Error
}

func (expenses *Expenses) FindAll() (*[]model.DailyExpenses, error) {
	var result *[]model.DailyExpenses

	if err := expenses.DB.Table("").Find(&result).Error; err != nil {
		return nil, errors.New("Error getting expenses")
	}
	return result, nil
}

func (expenses *Expenses) Update(exp *model.DailyExpenses) error {
	if exp == nil {
		errors.New("Error updating to expeses")
	}

	db := expenses.DB.Begin()

	if err := db.Table("").Save(&exp).Error; err != nil {
		return errors.New("Error updating to expenses")
	}

	return db.Commit().Error
}

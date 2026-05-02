package dal

import (
	"errors"
	model "tlaloc-catalog-service/model/db"

	"gorm.io/gorm"
)

type ExpensesCategoriesDAO interface {
	Create(ExpensesCategories *model.ExpensesCategories) error
	FindAll() ([]model.ExpensesCategories, error)
	Update(ExpesesCategories *model.ExpensesCategories) error
}

type ExpensesCategories struct {
	DB *gorm.DB
}

func NewExpensesCategories(db *gorm.DB) *ExpensesCategories {
	return &ExpensesCategories{
		DB: db,
	}
}

func (expensesCategories *ExpensesCategories) Create(ec *model.ExpensesCategories) error {
	if expensesCategories == nil {
		return errors.New("modelo vacio")
	}

	e := ec

	db := expensesCategories.DB.Begin()
	if err := db.Table("tlaloc_api.expense_categories").Select("name").Create(e).Error; err != nil {
		db.Rollback()
		return errors.New("creacion fallida")
	}

	return db.Commit().Error

}

func (expensesCategories *ExpensesCategories) FindAll() ([]model.ExpensesCategories, error) {
	var expensesCate []model.ExpensesCategories

	if err := expensesCategories.DB.Table("tlaloc_api.expense_categories").Find(&expensesCate).Error; err != nil {
		return nil, errors.New("error en la busqueda")
	}

	return expensesCate, nil
}

func (expensesCategories *ExpensesCategories) Update(ec *model.ExpensesCategories) error {
	if ec == nil {
		return errors.New("modelo vacio")
	}

	db := expensesCategories.DB.Begin()
	if error := db.Table("tlaloc_api.expense_categories").Save(ec).Error; error != nil {
		db.Rollback()
		return errors.New("error actualizando")
	}

	return db.Commit().Error
}

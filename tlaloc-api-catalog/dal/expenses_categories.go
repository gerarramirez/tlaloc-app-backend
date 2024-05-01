package dal

import (
	"errors"
	"time"
	model "tlaloc-catalog/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpensesCategoriesDAO interface {
	Create(ExpensesCategories *model.ExpensesCategories) error
	FindAll() (*[]model.ExpensesCategories, error)
	Update(ExpesesCategories *model.ExpensesCategories) error
}

type ExpensesCategories struct {
	DB           *gorm.DB
	GenerateUUID GenerateUUID
}

func NewExpensesCategories(db *gorm.DB) *ExpensesCategories {
	return &ExpensesCategories{
		DB:           db,
		GenerateUUID: func() string { return uuid.NewString() },
	}
}

func (expensesCategories *ExpensesCategories) Create(ec *model.ExpensesCategories) error {
	if expensesCategories == nil {
		return errors.New("modelo vacio")
	}

	e := &model.ExpensesCategoriesEntity{
		ExpensesCategories: model.ExpensesCategories{
			Name: ec.Name,
		},
		BaseEntity: model.BaseEntity{
			CreatedAt: time.Now(),
			ID:        expensesCategories.GenerateUUID(),
		},
	}

	db := expensesCategories.DB.Begin()
	if err := db.Table("").Create(e).Error; err != nil {
		db.Rollback()
		return errors.New("creacion fallida")
	}

	return db.Commit().Error

}

func (expensesCategories *ExpensesCategories) FindAll() ([]model.ExpensesCategories, error) {
	var expensesCate []model.ExpensesCategories

	if err := expensesCategories.DB.Table("").Find(expensesCate).Error; err != nil {
		return nil, errors.New("error en la busqueda")
	}

	return expensesCate, nil
}

func (expensesCategories *ExpensesCategories) Update(ec *model.ExpensesCategories) error {
	if ec == nil {
		return errors.New("modelo vacio")
	}

	db := expensesCategories.DB.Begin()
	if error := db.Table("").Save(ec).Error; error != nil {
		db.Rollback()
		return errors.New("error actualizando")
	}

	return db.Commit().Error
}

package dal

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	model "tlaloc-catalog/model/db"
)

type ExpensesDao interface {
	Create(expense *model.Expenses) error
	FindAll() (*[]model.Expenses, error)
	Update(expenses *model.Expenses) error
}

type Expenses struct {
	DB   *gorm.DB
	Uuid GenerateUUID
}

func NewExpensesDal(db *gorm.DB) *Expenses {
	return &Expenses{
		DB: db,
		Uuid: func() string {
			return uuid.NewString()
		},
	}
}

func (expenses *Expenses) Create(exp *model.Expenses) error {
	if exp == nil {
		return errors.New("modelo de expenses vacio")
	}

	model := &model.ExpensesEntity{
		Expenses: model.Expenses{
			Name: exp.Name,
		},
		BaseEntity: model.BaseEntity{
			ID:        expenses.Uuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db := expenses.DB.Begin()

	if err := db.Table("tlaloc_api.expenses").Create(model).Error; err != nil {
		db.Rollback()
		return errors.New("Error guardando expenses")
	}

	return db.Commit().Error
}

func (expenses *Expenses) FindAll() (*[]model.Expenses, error) {
	var result *[]model.Expenses

	if err := expenses.DB.Table("").Find(&result).Error; err != nil {
		return nil, errors.New("Error getting expenses")
	}
	return result, nil
}

func (expenses *Expenses) Update(exp *model.Expenses) error {
	if exp == nil {
		errors.New("Error updating to expeses")
	}

	db := expenses.DB.Begin()

	if err := db.Table("").Save(&exp).Error; err != nil {
		return errors.New("Error updating to expenses")
	}

	return db.Commit().Error
}

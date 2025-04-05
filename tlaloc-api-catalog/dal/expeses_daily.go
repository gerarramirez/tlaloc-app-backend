package dal

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	model "tlaloc-catalog/model/db"
)

type ExpensesDao interface {
	Create(expense *model.ExpensesDaily) error
	FindAll() (*[]model.ExpensesDaily, error)
	Update(expenses *model.ExpensesDaily) error
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

func (expenses *Expenses) Create(exp *model.ExpensesDaily) error {
	if exp == nil {
		return errors.New("modelo de expenses vacio")
	}

	model := &model.ExpensesEntity{
		ExpensesDaily: model.ExpensesDaily{
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

func (expenses *Expenses) FindAll() (*[]model.ExpensesDaily, error) {
	var result *[]model.ExpensesDaily

	if err := expenses.DB.Table("").Find(&result).Error; err != nil {
		return nil, errors.New("Error getting expenses")
	}
	return result, nil
}

func (expenses *Expenses) Update(exp *model.ExpensesDaily) error {
	if exp == nil {
		errors.New("Error updating to expeses")
	}

	db := expenses.DB.Begin()

	if err := db.Table("").Save(&exp).Error; err != nil {
		return errors.New("Error updating to expenses")
	}

	return db.Commit().Error
}

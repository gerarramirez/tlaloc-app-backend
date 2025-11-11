package dal

import (
	"gorm.io/gorm"
	"tlaloc-api-budgetManagement/model"
)

type BudgetDao interface {
	Create(budget *model.Budget) error
	Update(budget *model.Budget) error
	FindAll() ([]model.Budget, error)
}

type BudgetDaoImpl struct {
	DB *gorm.DB
}

func NewBudgetDal(db *gorm.DB) *BudgetDaoImpl {
	return &BudgetDaoImpl{
		DB: db,
	}
}

func (BudgetDaoImpl) Create(budget *model.Budget) error {
	//TODO implement me
	panic("implement me")
}

func (BudgetDaoImpl) Update(budget *model.Budget) error {
	//TODO implement me
	panic("implement me")
}

func (BudgetDaoImpl) FindAll() ([]model.Budget, error) {
	//TODO implement me
	panic("implement me")
}

package dal

import "tlaloc-api-budgetManagement/model"

type BudgetDao interface {
	Create(budget *model.Budget) error
	Update(budget *model.Budget) error
	FindAll() ([]model.Budget, error)
	FindByStartDate(budget *model.Budget) (model.Budget, error)
}

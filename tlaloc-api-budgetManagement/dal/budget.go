package dal

import (
	"errors"
	"gorm.io/gorm"
	"tlaloc-api-budgetManagement/model"
)

type BudgetDao interface {
	CreateBudget(budget *model.Budget) error

	CreateBudgetExpenseCate(budget *model.BudgetExpenseCategories) error

	GetWholeBudget(id *string) (*model.BudgetWithWholeExpensesCategory, error)
}

type BudgetDaoImpl struct {
	DB *gorm.DB
}

func NewBudgetDal(db *gorm.DB) *BudgetDaoImpl {
	return &BudgetDaoImpl{
		DB: db,
	}
}

func (dao *BudgetDaoImpl) CreateBudget(budget *model.Budget) error {

	db := dao.DB.Begin()

	if err := db.Select("assigned", "start_date", "end_date").Table("tlaloc_api.budgets").Create(&budget).Error; err != nil {
		db.Rollback()
		return errors.New("Error en el api  al hora de guardar")
	}
	//TODO implement me

	return db.Commit().Error
}

func (dao *BudgetDaoImpl) CreateBudgetExpenseCate(budget *model.BudgetExpenseCategories) error {

	db := dao.DB.Begin()

	if err := db.Table("tlaloc_api.budget_expense_categories").Select("expense_category_id", "budget_id", "assigned").Create(&budget).Error; err != nil {
		db.Rollback()
		return errors.New("Internal server error")
	}
	return db.Commit().Error
}

func (dao *BudgetDaoImpl) GetWholeBudget(id *string) (*model.BudgetWithWholeExpensesCategory, error) {
	budget := &model.Budget{}
	ExpensesCategories := &[]model.ExpensesCategories{}
	db := dao.DB.Begin()

	if err := db.Table("tlaloc_api.budgets").First(&budget, &id).Error; err != nil {
		return nil, errors.New("Id no econtrado")
	}

	db.Select("tlaloc_api.expense_categories.id as expense_category_id, tlaloc_api.expense_categories.name, tlaloc_api.budget_expense_categories.id as budget_expense_category_id, tlaloc_api.budget_expense_categories.assigned ").Table("tlaloc_api.budget_expense_categories").Joins("Inner Join tlaloc_api.expense_categories  on tlaloc_api.expense_categories.id = tlaloc_api.budget_expense_categories.expense_category_id", db.Where("tlaloc_api.budget_expense_categories.budget_id = ?", budget.Id)).Scan(&ExpensesCategories)

	return &model.BudgetWithWholeExpensesCategory{
		Budget:          budget,
		ExpenseCategory: ExpensesCategories,
	}, nil

}

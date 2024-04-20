package dal

import model "tlaloc-catalog/model/db"

type ExpensesCategoriesDAO interface {
	Create(ExpensesCategories model.ExpensesCategories) (*model.ExpensesCategories, error)
	FindAll() ([]model.ExpensesCategories, error)
	Update(ExpesesCategories model.ExpensesCategories) (*model.ExpensesCategories, error)
}

package model

import (
	"time"
)

type Budget struct {
	Id        string    `json:"id" param:"id"`
	Assigned  float32   `json:"assigned"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type BudgetExpenseCategories struct {
	Id                string  `json:"id"`
	ExpenseCategoryId string  `json:"expense_category_id"`
	BudgetId          string  `json:"budget_id"`
	Assigned          float32 `json:"assigned"`
}

type ExpensesCategories struct {
	ExpenseCategoryId       string  `json:"expense_category_id""`
	Name                    string  `json:"name"`
	BudgetExpenseCategoryId string  `json:"budget_expense_category_id"`
	Assigned                float32 `json:"assigned"`
}

type BudgetWithWholeExpensesCategory struct {
	Budget          *Budget               `json:"budget"`
	ExpenseCategory *[]ExpensesCategories `json:"budget_assigned_category"`
}

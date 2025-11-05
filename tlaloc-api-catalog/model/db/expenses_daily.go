package model

import "time"

type ExpensesDaily struct {
	Name                       string    `json:"name"`
	Amount                     float32   `json:"amount"`
	IdBudgetExpensesCategories string    `json:"id_fk_budget_expenses_categories"`
	IdCommerce                 string    `json:"id_commerce""`
	ExpensesDate               time.Time `json:"expenses_date"`
	IdBankProduct              string    `json:"id_bank_product""`
	Cash                       bool      `json:"cash"`
}

type ExpensesEntity struct {
	ExpensesDaily
	BaseEntity
}

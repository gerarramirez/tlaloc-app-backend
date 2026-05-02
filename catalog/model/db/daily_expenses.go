package model

import "time"

type DailyExpenses struct {
	Id                      string    `json:"id"`
	description             string    `json:"name"`
	Amount                  float32   `json:"amount"`
	BudgetExpenseCategoryId string    `json:"id_fk_budget_expenses_categories"`
	CommerceId              string    `json:"id_commerce""`
	ExpensesDate            time.Time `json:"expenses_date"`
	BankProductId           string    `json:"id_bank_product""`
}

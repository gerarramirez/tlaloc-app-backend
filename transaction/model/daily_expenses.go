package model

type DailyExpenses struct {
	Id                      string  `gorm:"primaryKey;default:null"json:"id"`
	Description             string  `json:"description"`
	BudgetExpenseCategoryId string  `json:"budget_expense_category_id"`
	CommerceId              string  `json:"commerce_id"`
	BankProductId           string  `json:"bank_product_id"`
	Amount                  float32 `json:"amount"`
}

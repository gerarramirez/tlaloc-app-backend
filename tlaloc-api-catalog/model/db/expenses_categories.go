package model

type ExpensesCategories struct {
	Name string `json:"name"`
}

type ExpensesCategoriesEntity struct {
	ExpensesCategories
	BaseEntity
}

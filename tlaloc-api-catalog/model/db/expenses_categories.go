package model

type ExpensesCategories struct {
	Id   string `json:"id"`
	name string `json:"name"`
}

type ExpensesCategoriesEntity struct {
	ExpensesCategories
	BaseEntity
}

package model

type ExpensesCategories struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ExpensesCategoriesEntity struct {
	ExpensesCategories
	BaseEntity
}

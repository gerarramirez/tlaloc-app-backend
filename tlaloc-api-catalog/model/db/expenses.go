package model

type Expenses struct {
	Name                 string `json: "name"`
	IdCategoriesExpenses string `json: "idCategoriesExpeses"`
}

type ExpensesEntity struct {
	Expenses
	BaseEntity
}

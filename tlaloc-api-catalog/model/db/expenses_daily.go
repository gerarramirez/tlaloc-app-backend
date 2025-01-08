package model

import "time"

type ExpensesDaily struct {
	Name string `json: "name"`
	//IdCategoriesExpenses     string    `json: "idCategoriesExpeses"`
	amount                   float32   `json:"amount"`
	idCommerce               string    `json:id_commerce`
	expensesDate             time.Time `json:expenses_date`
	transactionByBankProduct string    `json:id_bank_prodcut`
	cash                     bool      `json:"cash"`
}

type ExpensesEntity struct {
	ExpensesDaily
	BaseEntity
}

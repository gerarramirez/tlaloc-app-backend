package model

type Bank struct {
	Name string `json:"name"`
}

type BankEntity struct {
	BaseEntity
	Bank
}

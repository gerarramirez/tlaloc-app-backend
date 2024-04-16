package model

type BankJson struct {
	Name string `json:"name"`
}

type Bank struct {
	BaseEntity
	BankJson
}

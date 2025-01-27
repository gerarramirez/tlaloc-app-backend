package model

type BanksProducts struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	IdBank       string  `json:"id_bank"`
	Amount       float32 `json:"amount"`
	Card         bool    `json:"card"`
	Account      bool    `json:"account"`
	Loan         bool    `json:"loan"`
	InterestRate float32 `json:"interest_rate"`
}

type BanksProductsEntity struct {
	BanksProducts
	BaseEntity
}

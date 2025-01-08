package model

type BanksProducts struct {
	Name         string  `json:"name"`
	IdBank       string  `json:"id_bank"`
	amount       float32 `json:"amount"`
	card         bool    `json:"card"`
	account      bool    `json:"account"`
	loan         bool    `json:"loan"`
	interestRate float32 `json:"interest_rate"`
}

type BanksProductsEntity struct {
	BanksProducts
	BaseEntity
}

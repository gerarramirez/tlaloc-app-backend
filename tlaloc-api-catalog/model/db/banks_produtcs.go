package model

type BanksProducts struct {
	Name   string `json:"name"`
	IdBank string `json:"id_bank"`
}

type BanksProductsEntity struct {
	BanksProducts
	BaseEntity
}

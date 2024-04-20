package model

type BanksProducts struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	IdBank string `json:"id_bank"`
}

type BanksProductsEntity struct {
	BanksProducts
	BaseEntity
}

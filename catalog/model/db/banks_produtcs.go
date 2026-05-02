package model

type BanksProducts struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	AccountNumber  string `json:"account-number"`
	IdBank         string `json:"id_bank"`
	ProductType    string `json:"id_product_type"`
	InterestRateId string `json:"interest_rate"`
}

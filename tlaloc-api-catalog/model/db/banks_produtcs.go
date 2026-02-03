package model

type BanksProducts struct {
	Id             string `gorm:"primaryKey;default:null" json:"id"`
	Name           string `json:"name"`
	AccountNumber  string `json:"account_number"`
	BankId         string `json:"bank_id"`
	ProductTypeId  string `json:"product_type_id"`
	InterestRateId string `json:"interest_rate_id"`
}

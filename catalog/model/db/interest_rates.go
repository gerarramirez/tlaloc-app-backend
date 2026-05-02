package model

type InterestRate struct {
	Id      string  `gorm:"primaryKey;default:null" json:"id"`
	Percent float32 `json:"percent"`
	Type    string  `json:"type"`
}

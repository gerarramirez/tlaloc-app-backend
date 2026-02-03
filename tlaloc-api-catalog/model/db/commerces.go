package model

type Commerces struct {
	Id                 string `gorm:"primaryKey;default:null"json:"id"`
	Name               string `json:"name"`
	CommerceCategoryId string `json:"commerce_category_id"`
}

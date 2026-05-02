package model

type CommercesSubcategories struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	CommerceCategoryId string `json:"id_commerce_category"`
}

package model

type CommercesCategories struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CommercesCategoriesEntity struct {
	CommercesCategories
	BaseEntity
}

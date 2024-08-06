package model

type CommercesCategories struct {
	Name string `json:"name"`
}

type CommercesCategoriesEntity struct {
	CommercesCategories
	BaseEntity
}

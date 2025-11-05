package model

type CommercesSubcategories struct {
	Name                  string `json:"name"`
	IdCommercesCategories string `json:"id_commerces_categories"`
}

type CommercesSubcategoriesEntity struct {
	CommercesSubcategories
	BaseEntity
}

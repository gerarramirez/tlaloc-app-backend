package model

type CommercesSubcategories struct {
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	IdCommercesCategories string `json:"id_commerces_categories"`
}

type CommercesSubcategoriesEntity struct {
	CommercesSubcategories
	BaseEntity
}

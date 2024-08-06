package model

type Commerces struct {
	Name                  string `json:"name"`
	IdCommercesCategories string `json:"id_commerces_categories"`
}

type CommercesEntity struct {
	Commerces
	BaseEntity
}

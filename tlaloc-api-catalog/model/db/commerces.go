package model

type Commerces struct {
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	IdCommercesCategories string `json:"id_commerces_categories"`
}

type CommercesEntity struct {
	Commerces
	BaseEntity
}

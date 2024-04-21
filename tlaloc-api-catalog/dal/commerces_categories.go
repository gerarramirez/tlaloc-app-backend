package dal

import model "tlaloc-catalog/model/db"

type CommercesCategoriesDAO interface {
	Create(CommercesCategories *model.CommercesCategories) (*model.CommercesCategories, error)
	FindAll() ([]model.CommercesCategories, error)
	Update(CommercesCategories *model.CommercesCategories) (*model.CommercesCategories, error)
}

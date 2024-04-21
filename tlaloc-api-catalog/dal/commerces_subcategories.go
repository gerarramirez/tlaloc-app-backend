package dal

import model "tlaloc-catalog/model/db"

type CommercesSubcategoriesDAO interface {
	Create(CommercesSubcategories *model.CommercesSubcategories) (*model.CommercesSubcategories, error)
	FindAll() ([]model.CommercesSubcategories, error)
	Update(CommercesSubcategories *model.CommercesSubcategories) (*model.CommercesSubcategories, error)
}

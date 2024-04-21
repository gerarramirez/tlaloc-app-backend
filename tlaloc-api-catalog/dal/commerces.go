package dal

import model "tlaloc-catalog/model/db"

type CommercesDAO interface {
	Create(Commerces *model.Commerces) (*model.Commerces, error)
	FindAll() (*[]model.Commerces, error)
	Update(Commerce *model.Commerces) (*model.Commerces, error)
}

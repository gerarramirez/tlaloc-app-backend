package dal

import model "tlaloc-catalog/model/db"

type BanksProductDao interface {
	Create(banksProducts *model.BanksProducts) (*model.BanksProducts, error)
	FindAll() ([]model.BanksProducts, error)
	Update(banksProducts *model.BanksProducts) (*model.BanksProducts, error)
}

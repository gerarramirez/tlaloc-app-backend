package dal

import (
	"errors"
	model "tlaloc-catalog-service/model/db"

	"gorm.io/gorm"
)

type CommercesDAO interface {
	Create(Commerces *model.Commerces) (*model.Commerces, error)
	FindAll() ([]model.Commerces, error)
	Update(Commerce *model.Commerces) (*model.Commerces, error)
}

type Commerces struct {
	DB *gorm.DB
}

func NewCommercesDal(db *gorm.DB) *Commerces {
	return &Commerces{
		DB: db,
	}
}

func (comm *Commerces) Create(commerces *model.Commerces) (*model.Commerces, error) {

	if commerces == nil {
		return nil, errors.New("modelo vacio")
	}

	c := commerces

	db := comm.DB.Begin()
	if err := db.Table("tlaloc_api.commerces").Create(&c).Error; err != nil {
		db.Rollback()
		return nil, err
	}

	db.Commit()

	return commerces, nil

}

func (comm *Commerces) FindAll() ([]model.Commerces, error) {

	var (
		commerces []model.Commerces
	)

	if error := comm.DB.Table("tlaloc_api.commerces").Find(&commerces).Error; error != nil {
		return nil, errors.New("error en la extraccion de los commercios")
	}

	return commerces, nil
}

func (comm *Commerces) Update(c *model.Commerces) (*model.Commerces, error) {
	if c == nil {
		return nil, errors.New("modelo vacio")
	}

	db := comm.DB.Begin()
	if err := db.Table("tlaloc_api.commerces").Save(c).Error; err != nil {
		return nil, errors.New("error en el actaulizado de comercios")
	}

	db.Commit()

	return c, nil

}

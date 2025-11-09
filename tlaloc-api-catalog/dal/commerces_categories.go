package dal

import (
	"errors"
	model "tlaloc-catalog/model/db"

	"gorm.io/gorm"
)

type CommercesCategoriesDAO interface {
	Create(CommercesCategories *model.CommercesCategories) error
	FindAll() ([]model.CommercesCategories, error)
	Update(CommercesCategories *model.CommercesCategories) error
}

type CommercesCategories struct {
	DB *gorm.DB
}

func NewCommercesCategories(db *gorm.DB) *CommercesCategories {
	return &CommercesCategories{
		DB: db,
	}
}

func (commercesDao CommercesCategories) Create(com *model.CommercesCategories) error {
	if com == nil {
		errors.New("error en el modelo de modelo de comercio")
	}

	c := com

	db := commercesDao.DB.Begin()
	if err := db.Table("tlaloc_api.commerces_categories").Create(c).Error; err != nil {
		db.Rollback()
		return errors.New("error en el guardado de la tabla")
	}

	return db.Commit().Error
}

func (commercesCategoriesDao CommercesCategories) FindAll() ([]model.CommercesCategories, error) {

	var commercesCategories []model.CommercesCategories

	if err := commercesCategoriesDao.DB.Table("tlaloc_api.commerces_categories").Find(&commercesCategories).Error; err != nil {
		return nil, errors.New("error en la extraccion de datos")
	}

	return commercesCategories, nil

}

func (commercesCategoriesDao CommercesCategories) Update(commercesCategories *model.CommercesCategories) error {
	if commercesCategories == nil {
		return errors.New("modelo vacio")
	}

	c := commercesCategories

	db := commercesCategoriesDao.DB.Begin()

	if err := db.Table("tlaloc_api.commerces_categories").Save(c).Error; err != nil {
		db.Rollback()
		print("error ")
		return errors.New("no se pudo actualizar el registro")
	}

	return db.Commit().Error
}

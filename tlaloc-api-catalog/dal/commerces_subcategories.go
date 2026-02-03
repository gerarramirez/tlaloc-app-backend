package dal

import (
	"errors"
	model "tlaloc-catalog/model/db"

	"gorm.io/gorm"
)

type CommercesSubcategoriesDAO interface {
	Create(CommercesSubcategories *model.CommercesSubcategories) error
	FindAll() ([]model.CommercesSubcategories, error)
	Update(CommercesSubcategories *model.CommercesSubcategories) error
}

type CommercesSubcategories struct {
	DB *gorm.DB
}

func NewCommercesSubcategories(db *gorm.DB) *CommercesSubcategories {
	return &CommercesSubcategories{
		DB: db,
	}
}

func (cs *CommercesSubcategories) Create(commercesSubcategories *model.CommercesSubcategories) error {
	if commercesSubcategories == nil {
		return errors.New("modelo vacio")
	}

	c := commercesSubcategories

	db := cs.DB.Begin()

	if err := db.Select("name", "commerce_category_id").Table("tlaloc_api.commerce_subcategories").Create(c).Error; err != nil {
		db.Rollback()
		return errors.New("error en el guardo de la persistencia")
	}

	return db.Commit().Error
}

func (CommercesSubCategoriesDao *CommercesSubcategories) FindAll() ([]model.CommercesSubcategories, error) {

	var commercesSubcategories []model.CommercesSubcategories

	if err := CommercesSubCategoriesDao.DB.Table("tlaloc_api.commerce_subcategories").Find(&commercesSubcategories).Error; err != nil {
		return nil, errors.New("error en el modelo")

	}

	return commercesSubcategories, nil

}

func (CommercesSubcategoriesDAO *CommercesSubcategories) Update(commercesSubcategories *model.CommercesSubcategories) error {
	if commercesSubcategories == nil {
		return errors.New("modelo vacio")
	}

	cs := commercesSubcategories

	db := CommercesSubcategoriesDAO.DB.Begin()

	if err := db.Select("name", "commerce_category_id").Table("tlaloc_api.commerce_subcategories").Save(&cs).Error; err != nil {
		db.Rollback()
		return errors.New("error en el guardado")
	}

	return db.Commit().Error
}

package dal

import (
	"errors"
	"time"
	model "tlaloc-catalog/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommercesSubcategoriesDAO interface {
	Create(CommercesSubcategories *model.CommercesSubcategories) error
	FindAll() ([]model.CommercesSubcategories, error)
	Update(CommercesSubcategories *model.CommercesSubcategories) error
}

type CommercesSubcategories struct {
	DB           *gorm.DB
	GenerateUUID GenerateUUID
}

func NewCommercesSubcategories(db *gorm.DB) *CommercesSubcategories {
	return &CommercesSubcategories{
		DB: db,
		GenerateUUID: func() string {
			return uuid.NewString()
		},
	}
}

func (cs *CommercesSubcategories) Create(commercesSubcategories *model.CommercesSubcategories) error {
	if commercesSubcategories == nil {
		return errors.New("modelo vacio")
	}

	c := &model.CommercesSubcategoriesEntity{
		CommercesSubcategories: model.CommercesSubcategories{
			Name:                  commercesSubcategories.Name,
			IdCommercesCategories: commercesSubcategories.IdCommercesCategories,
		},
		BaseEntity: model.BaseEntity{
			ID:        cs.GenerateUUID(),
			CreatedAt: time.Now(),
		},
	}

	db := cs.DB.Begin()

	if err := db.Table("").Create(c).Error; err != nil {
		db.Rollback()
		return errors.New("error en el guardo de la persistencia")
	}

	return db.Commit().Error
}

func (CommercesSubCategoriesDao *CommercesSubcategories) FindAll() ([]model.CommercesSubcategories, error) {

	var commercesSubcategories []model.CommercesSubcategories

	if err := CommercesSubCategoriesDao.DB.Table("").Find(&commercesSubcategories).Error; err != nil {
		return nil, errors.New("error en el modelo")

	}

	return commercesSubcategories, nil

}

func (CommercesSubcategoriesDAO *CommercesSubcategories) Update(commercesSubcategories *model.CommercesSubcategories) error {
	if commercesSubcategories == nil {
		return errors.New("modelo vacio")
	}

	cs := &model.CommercesSubcategoriesEntity{
		CommercesSubcategories: model.CommercesSubcategories{
			Id:                    commercesSubcategories.Id,
			Name:                  commercesSubcategories.Name,
			IdCommercesCategories: commercesSubcategories.IdCommercesCategories,
		},
		BaseEntity: model.BaseEntity{
			UpdatedAt: time.Now(),
		},
	}

	db := CommercesSubcategoriesDAO.DB.Begin()

	if err := db.Table("").Save(&cs).Error; err != nil {
		db.Rollback()
		return errors.New("error en el guardado")
	}

	return db.Commit().Error
}

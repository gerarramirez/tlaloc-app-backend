package dal

import (
	"errors"
	"time"
	model "tlaloc-catalog/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommercesDAO interface {
	Create(Commerces *model.Commerces) (*model.Commerces, error)
	FindAll() ([]model.Commerces, error)
	Update(Commerce *model.Commerces) (*model.Commerces, error)
}

type Commerces struct {
	DB   *gorm.DB
	uuid GenerateUUID
}

func NewCommercesDal(db *gorm.DB) *Commerces {
	return &Commerces{
		DB:   db,
		uuid: func() string { return uuid.New().String() },
	}
}

func (comm *Commerces) Create(commerces *model.Commerces) (*model.Commerces, error) {

	if commerces == nil {
		return nil, errors.New("modelo vacio")
	}

	now := time.Now()

	c := &model.CommercesEntity{
		Commerces: model.Commerces{
			Name:                  commerces.Name,
			IdCommercesCategories: commerces.IdCommercesCategories,
		},
		BaseEntity: model.BaseEntity{
			ID:        comm.uuid(),
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	db := comm.DB.Begin()
	if err := db.Table("").Create(&c).Error; err != nil {
		db.Rollback()
		return nil, err
	}

	db.Commit()

	commerces.Id = c.BaseEntity.ID

	return commerces, nil

}

func (comm *Commerces) FindAll() ([]model.Commerces, error) {

	var (
		commerces []model.Commerces
	)

	if error := comm.DB.Table("").Find(&commerces).Error; error != nil {
		return nil, errors.New("error en la extraccion de los commercios")
	}

	return commerces, nil
}

func (comm *Commerces) Update(c *model.Commerces) (*model.Commerces, error) {
	if c == nil {
		return nil, errors.New("modelo vacio")
	}

	db := comm.DB.Begin()
	if err := db.Table("").Save(c).Error; err != nil {
		return nil, errors.New("error en el actaulizado de comercios")
	}

	db.Commit()

	return c, nil

}

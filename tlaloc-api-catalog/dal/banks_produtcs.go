package dal

import (
	"errors"
	"time"
	model "tlaloc-catalog/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BanksProductDao interface {
	Create(banksProducts *model.BanksProducts) error
	FindAll() ([]model.BanksProducts, error)
	Update(banksProducts *model.BanksProducts) error
}

type BanksProducts struct {
	DB           *gorm.DB
	GenerateUUID GenerateUUID
}

func NewBanksProducts(db *gorm.DB) *BanksProducts {
	return &BanksProducts{
		DB: db,
		GenerateUUID: func() string {
			return uuid.New().String()
		},
	}
}

func (b *BanksProducts) Create(bankProducts *model.BanksProducts) error {
	if bankProducts == nil {
		return errors.New("productos bancarios viene vacio")
	}

	bpEntity := &model.BanksProductsEntity{
		BanksProducts: model.BanksProducts{
			Name:   bankProducts.Name,
			IdBank: bankProducts.IdBank,
		},
		BaseEntity: model.BaseEntity{
			ID:        b.GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db := b.DB.Begin()
	if err := db.Table("").Create(&bpEntity).Error; err != nil {
		db.Rollback()
		return errors.New("error en el guardado de la informacion de producto bancario")
	}
	return db.Commit().Error
}

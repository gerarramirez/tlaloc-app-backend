package dal

import (
	"errors"
	"time"
	model "tlaloc-catalog/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const uuidLength = 36

type BankDAO interface {
	Create(bank *model.Bank) (*model.BankEntity, error)
}

type Bank struct {
	DB           *gorm.DB
	GenerateUUID GenerateUUID
}

func NewBankDal(db *gorm.DB) *Bank {
	return &Bank{
		DB: db,
		GenerateUUID: func() string {
			return uuid.New().String()
		},
	}
}

func (b *Bank) Create(bank *model.Bank) (*model.BankEntity, error) {
	if bank == nil {
		return nil, errors.New("Bank can not be nil")
	}

	println("esta es lo que trae " + bank.Name)
	now := time.Now()

	e := &model.BankEntity{
		BaseEntity: model.BaseEntity{
			ID:        b.GenerateUUID(),
			CreatedAt: now,
			UpdatedAt: now,
		},
		Bank: model.Bank{
			Name: bank.Name,
		},
	}

	db := b.DB.Begin()
	if err := db.Table("tlaloc_api.banks").Create(&e).Error; err != nil {
		println("error garrafal perro " + err.Error())
		db.Rollback()
		return e, err
	}

	db.Commit()

	return e, nil

}

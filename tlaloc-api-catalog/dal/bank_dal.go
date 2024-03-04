package dal

import (
	"context"
	"errors"
	"time"
	model "tlaloc-catalog/model/db"

	"gorm.io/gorm"
)

const uuidLength = 36

type Bank struct {
	DB           *gorm.DB
	GenerateUUID GenerateUUID
}

func (b *Bank) Create(ctx context.Context, bank *model.Bank) (*model.BankEntity, error) {
	if bank == nil {
		return nil, errors.New("Bank can not be nil")
	}

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
	if err := db.Create(&e).Error; err != nil {
		db.Rollback()
		return e, err
	}

	return e, nil

}

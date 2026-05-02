package dal

import (
	"errors"
	model "tlaloc-catalog-service/model/db"

	"gorm.io/gorm"
)

type BankDAO interface {
	Create(bank *model.Bank) error
	FindAll() ([]model.Bank, error)
	Update(ba *model.Bank) error
}

type Bank struct {
	DB *gorm.DB
}

func NewBankDal(db *gorm.DB) *Bank {
	return &Bank{
		DB: db,
	}
}

func (b *Bank) Create(bank *model.Bank) error {
	if bank == nil {
		return errors.New("Bank can not be nil")
	}

	println("esta es lo que trae " + bank.Name)

	e := &model.Bank{
		Name: bank.Name,
	}

	db := b.DB.Begin()
	if err := db.Select("Name").Table("tlaloc_api.banks").Create(&e).Error; err != nil {
		db.Rollback()
		return err
	}

	db.Commit()

	return nil

}

func (b *Bank) FindAll() ([]model.Bank, error) {
	var (
		banks []model.Bank
	)

	if err := b.DB.Table("tlaloc_api.banks").Find(&banks).Error; err != nil {
		return nil, err
	}

	return banks, nil
}

func (b *Bank) Update(ba *model.Bank) error {
	db := b.DB.Begin()
	if err := db.Table("tlaloc_api.banks").Save(&ba).Error; err != nil {
		db.Rollback()
		return err
	}

	return db.Commit().Error
}

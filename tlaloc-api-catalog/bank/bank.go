package bank

import (
	model "tlaloc-catalog/model/db"
)

type Bank_dal interface {
	Create(bank *model.BankJson) (*model.Bank, error)
}

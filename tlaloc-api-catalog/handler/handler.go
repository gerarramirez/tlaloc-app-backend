package handler

import (
	"tlaloc-catalog/bank"
)

type Handler struct {
	bankDal bank.Bank_dal
}

func NewHandler(bank bank.Bank_dal) *Handler {
	return &Handler{bankDal: bank}
}

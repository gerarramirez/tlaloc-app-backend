package handler

import (
	"tlaloc-catalog/dal"
)

type Handler struct {
	bankDAO dal.BankDAO
}

func NewHandler(bank dal.BankDAO) *Handler {
	return &Handler{bankDAO: bank}
}

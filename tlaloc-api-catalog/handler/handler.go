package handler

import (
	"tlaloc-catalog/dal"
)

type Handler struct {
	bankDAO                   dal.BankDAO
	banksProductsDAO          dal.BanksProductDao
	commercesDAO              dal.CommercesDAO
	commercesCategoriesDAO    dal.CommercesCategoriesDAO
	commercesSubcategoriesDAO dal.CommercesSubcategoriesDAO
	expensesCategoriesDAO     dal.ExpensesCategoriesDAO
}

func NewHandler(bank dal.BankDAO) *Handler {
	return &Handler{bankDAO: bank}
}

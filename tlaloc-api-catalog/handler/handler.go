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
	expenses                  dal.ExpensesDao
	incomeTypesDAO            dal.IncomeTypeDAO
}

func NewHandler(bank dal.BankDAO, bankProducts dal.BanksProductDao, commerceCategories dal.CommercesCategoriesDAO, commercesSubcategories dal.CommercesSubcategoriesDAO, commerces dal.CommercesDAO, expensesCategories dal.ExpensesCategoriesDAO, expenses dal.ExpensesDao,
	incomeTypes dal.IncomeTypeDAO) *Handler {
	return &Handler{
		bankDAO:                   bank,
		banksProductsDAO:          bankProducts,
		commercesCategoriesDAO:    commerceCategories,
		commercesSubcategoriesDAO: commercesSubcategories,
		commercesDAO:              commerces,
		expensesCategoriesDAO:     expensesCategories,
		expenses:                  expenses,
		incomeTypesDAO:            incomeTypes,
	}
}

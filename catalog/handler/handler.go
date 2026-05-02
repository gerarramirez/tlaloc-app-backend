package handler

import (
	"tlaloc-catalog-service/dal"
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
	productTypeDAO            dal.ProductTypeDAO
	interestRateDao           dal.InterestRateDAO
}

func NewHandler(bank dal.BankDAO, bankProducts dal.BanksProductDao, commerceCategories dal.CommercesCategoriesDAO, comerciosSubcategories dal.CommercesSubcategoriesDAO, comercios dal.CommercesDAO, expensesCategories dal.ExpensesCategoriesDAO, expenses dal.ExpensesDao,
	incomeTypes dal.IncomeTypeDAO, productTypeDAO dal.ProductTypeDAO, interestRate dal.InterestRateDAO) *Handler {
	return &Handler{
		bankDAO:                   bank,
		banksProductsDAO:          bankProducts,
		commercesCategoriesDAO:    commerceCategories,
		commercesSubcategoriesDAO: comerciosSubcategories,
		commercesDAO:              comercios,
		expensesCategoriesDAO:     expensesCategories,
		expenses:                  expenses,
		incomeTypesDAO:            incomeTypes,
		productTypeDAO:            productTypeDAO,
		interestRateDao:           interestRate,
	}
}

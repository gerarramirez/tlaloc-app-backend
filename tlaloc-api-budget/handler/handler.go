package handler

import "tlaloc-api-budget/dal"

type Handler struct {
	budgetDao dal.BudgetDao
}

func NewHandler(budget dal.BudgetDao) *Handler {
	return &Handler{
		budgetDao: budget,
	}
}

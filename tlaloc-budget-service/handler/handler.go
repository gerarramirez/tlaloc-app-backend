package handler

import "tlaloc-budget-service/dal"

type Handler struct {
	budgetDao dal.BudgetDao
}

func NewHandler(budget dal.BudgetDao) *Handler {
	return &Handler{
		budgetDao: budget,
	}
}

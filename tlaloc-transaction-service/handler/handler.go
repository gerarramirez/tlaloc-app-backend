package handler

import (
	"tlaloc-transaction-service/dal"
)

type Handler struct {
	dailyExpensesDao dal.DailyExpensesDao
}

func NewHandler(dao dal.DailyExpensesDao) *Handler {
	return &Handler{
		dailyExpensesDao: dao,
	}
}

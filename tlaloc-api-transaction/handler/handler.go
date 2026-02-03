package handler

import (
	"tlaloc-api-transaction/dal"
)

type Handler struct {
	dailyExpensesDao dal.DailyExpensesDao
}

func NewHandler(dao dal.DailyExpensesDao) *Handler {
	return &Handler{
		dailyExpensesDao: dao,
	}
}

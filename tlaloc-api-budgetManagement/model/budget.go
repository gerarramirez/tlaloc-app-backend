package model

import "time"

type Budget struct {
	Assigned  float32   `json:"assigned"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type BudgetEntity struct {
	Budget
	BaseEntity
}

package model

import (
	"time"
)

type BaseEntity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

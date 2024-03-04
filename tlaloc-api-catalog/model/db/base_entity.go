package model

import (
	"database/sql"
	"time"
)

type BaseEntity struct {
	ID        string       `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeleteAt  sql.NullTime `json:"deleted_at"`
}

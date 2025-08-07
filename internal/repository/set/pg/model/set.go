package model

import (
	"database/sql"
	"time"
)

type Set struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type NewSet struct {
	Name string
}

type UpdatedSet struct {
	Name sql.NullString
}

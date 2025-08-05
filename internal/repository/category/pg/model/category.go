package model

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int64
	Name      string
	DomainID  int64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type NewCategory struct {
	Name     string
	DomainID int64
}

type UpdatedCategory struct {
	Name     sql.NullString
	DomainID sql.NullInt64
}

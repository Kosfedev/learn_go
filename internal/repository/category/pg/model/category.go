package model

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int32
	Name      string
	DomainID  int32
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type NewCategory struct {
	Name     string
	DomainID int32
}

type UpdatedCategory struct {
	Name     sql.NullString
	DomainID sql.NullInt32
}

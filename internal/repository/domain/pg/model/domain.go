package model

import (
	"database/sql"
	"time"
)

type Domain struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type NewDomain struct {
	Name string
}

type UpdatedDomain struct {
	Name sql.NullString
}

package model

import (
	"database/sql"
	"time"
)

type Category struct {
	Id        int32
	Name      string
	DomainId  int32
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type NewCategory struct {
	Name     string
	DomainId int32
}

type UpdatedCategory struct {
	Name     sql.NullString
	DomainId sql.NullInt32
}

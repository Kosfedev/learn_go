package model

import "time"

type Category struct {
	Id        int
	Name      string
	DomainId  int
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type NewCategory struct {
	Name     string
	DomainId int
}

type UpdatedCategory struct {
	Name     *string
	DomainId *int
}

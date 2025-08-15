package model

import "time"

type Set struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type NewSet struct {
	Name string
}

type UpdatedSet struct {
	Name *string
}

type SetListSearchOptions struct {
	Filters *struct {
		Name string
	}
	Pagination *Pagination
	Sort       *Sort
}

type SetListWithTotal struct {
	Sets  []*Set
	Total int64
}

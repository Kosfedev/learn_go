package model

const (
	ASC = iota
	DESC
)

type Pagination struct {
	Offset int64
	Limit  int64
}

type PaginationWithTotal struct {
	Offset int64
	Limit  int64
	Total  int64
}

type Sort struct {
	SortBy    string
	SortOrder SortOrder
}

type SortOrder uint8

func (sortOrder SortOrder) IsValid() bool {
	if sortOrder == ASC || sortOrder == DESC {
		return true
	}

	return false
}

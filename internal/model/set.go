package model

import "time"

type Set struct {
	ID          int64
	Name        string
	QuestionIDs []int64
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type NewSet struct {
	Name        string
	QuestionIDs []int64
}

type UpdatedSet struct {
	Name *string
}

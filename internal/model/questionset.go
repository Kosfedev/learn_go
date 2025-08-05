package model

import "time"

type QuestionSet struct {
	ID          int64
	Name        string
	QuestionIDs []int64
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type NewQuestionSet struct {
	Name        string
	QuestionIDs []int64
}

type UpdatedQuestionSet struct {
	Name *string
}

package model

import "time"

type QuestionsDomain struct {
	Id   int
	Name string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type NewQuestionsDomain struct {
	Name string
}

type UpdatedQuestionsDomain struct {
	Name *string
}

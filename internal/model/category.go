package model

import "time"

type QuestionsCategory struct {
	Id   int
	Name string
	DomainId int
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type NewQuestionsCategory struct {
	Name string
	DomainId int
}

type UpdatedQuestionsCategory struct {
	Name *string
	DomainId *int
}

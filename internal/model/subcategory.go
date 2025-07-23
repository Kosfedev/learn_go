package model

import "time"

type QuestionsSubcategory struct {
	Id         int
	Name       string
	CategoryId int
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

type NewQuestionsSubcategory struct {
	Name       string
	CategoryId int
}

type UpdatedQuestionsSubcategory struct {
	Name       *string
	CategoryId *int
}

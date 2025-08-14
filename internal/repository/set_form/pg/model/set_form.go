package model

import (
	modelQuestionRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
	modelSetRepo "github.com/Kosfedev/learn_go/internal/repository/set/pg/model"
)

type SetForm struct {
	Set       *modelSetRepo.Set             `db:"set" json:"set"`
	Questions []*modelQuestionRepo.Question `db:"questions" json:"questions"`
}

package model

import (
	modelQuestionRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
	modelQuestionOptionRepo "github.com/Kosfedev/learn_go/internal/repository/question_option/pg/model"
	modelSetRepo "github.com/Kosfedev/learn_go/internal/repository/set/pg/model"
	modelSubcategoryRepo "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/model"
)

type QuestionForm struct {
	Question      *modelQuestionRepo.Question               `db:"question" json:"question"`
	Options       []*modelQuestionOptionRepo.QuestionOption `db:"question_options" json:"question_options"`
	Sets          []*modelSetRepo.Set                       `db:"sets" json:"sets"`
	Subcategories []*modelSubcategoryRepo.Subcategory       `db:"subcategories" json:"subcategories"`
}

type QuestionWithOptions struct {
	Question *modelQuestionRepo.Question               `db:"question" json:"question"`
	Options  []*modelQuestionOptionRepo.QuestionOption `db:"question_options" json:"question_options"`
}

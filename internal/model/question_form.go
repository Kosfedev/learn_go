package model

type QuestionForm struct {
	Question      *Question
	Options       []*QuestionOption
	Sets          []*Set
	Subcategories []*Subcategory
}

type QuestionWithOptions struct {
	Question *Question
	Options  []*QuestionOption
}

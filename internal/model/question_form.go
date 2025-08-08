package model

type QuestionForm struct {
	Question      *Question
	Options       []*QuestionOption
	Sets          []*Set
	Subcategories []*Subcategory
}

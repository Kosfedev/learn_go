package converter

import (
	"github.com/Kosfedev/learn_go/internal/model"
	converterQuestion "github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_form/pg/model"
	"github.com/Kosfedev/learn_go/internal/repository/question_option/pg/converter"
	converterSet "github.com/Kosfedev/learn_go/internal/repository/set/pg/converter"
	converterSubcategory "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/converter"
)

func QuestionFormFromPGSQL(questionFormRepo *modelRepo.QuestionForm) *model.QuestionForm {
	questionForm := &model.QuestionForm{
		Question:      converterQuestion.QuestionFromPGSQL(questionFormRepo.Question),
		Options:       converter.QuestionOptionsFromPGSQL(questionFormRepo.Options),
		Sets:          converterSet.SetsFromPGSQL(questionFormRepo.Sets),
		Subcategories: converterSubcategory.SubcategoriesFromPGSQL(questionFormRepo.Subcategories),
	}

	return questionForm
}

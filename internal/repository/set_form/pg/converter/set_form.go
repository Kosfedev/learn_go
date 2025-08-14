package converter

import (
	"github.com/Kosfedev/learn_go/internal/model"
	converterQuestion "github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
	converterSet "github.com/Kosfedev/learn_go/internal/repository/set/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/set_form/pg/model"
)

func SetFormFromPGSQL(setFormRepo *modelRepo.SetForm) *model.SetForm {
	setForm := &model.SetForm{
		Set:       converterSet.SetFromPGSQL(setFormRepo.Set),
		Questions: converterQuestion.QuestionsFromPGSQL(setFormRepo.Questions),
	}

	return setForm
}

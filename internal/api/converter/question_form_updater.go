package converter

import (
	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func NewQuestionWithOptionsFromGRPC(req *desc.CreateWithOptionsRequest) *model.NewQuestion {
	newOptions := make([]*model.NewQuestionOption, len(req.Data.Options))
	for i, optionGRPC := range req.Data.Options {
		optionServ := &model.NewQuestionOption{
			Text:      optionGRPC.Text,
			IsCorrect: optionGRPC.IsCorrect,
		}

		newOptions[i] = optionServ
	}

	newQuestion := NewQuestionFromGRPC(req.Data.Question)
	newQuestion.Options = newOptions

	return newQuestion
}

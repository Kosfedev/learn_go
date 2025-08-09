package converter

import (
	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func NewQuestionWithOptionsFromGRPC(req *desc.CreateWithOptionsRequest) *model.NewQuestionWithOptions {
	newQuestionWithOptions := &model.NewQuestionWithOptions{
		Question: NewQuestionFromGRPC(req.Data.Question),
		Options:  NewQuestionOptionsFromGRPC(req.Data.Options),
	}

	return newQuestionWithOptions
}

func NewQuestionOptionsFromGRPC(newOptionsGRPC []*desc.NewQuestionOption) []*model.NewQuestionOption {
	newOptions := make([]*model.NewQuestionOption, len(newOptionsGRPC))
	for i, option := range newOptionsGRPC {
		newOptions[i] = &model.NewQuestionOption{
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		}
	}

	return newOptions
}

package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func NewQuestionFromGRPC(req *desc.CreateRequest) *model.NewQuestion {
	newQuestion := &model.NewQuestion{
		Text: req.Text,
		Type: model.QuestionType(req.QuestionType),
	}

	if req.ReferenceAnswer != nil {
		refAnswer := req.ReferenceAnswer.GetValue()
		newQuestion.ReferenceAnswer = &refAnswer
	}

	return newQuestion
}

func UpdatedQuestionFromGRPC(req *desc.UpdateRequest) *model.UpdatedQuestion {
	updatedQuestion := &model.UpdatedQuestion{}

	if req.ReferenceAnswer != nil {
		refAnswer := req.ReferenceAnswer.GetValue()
		updatedQuestion.ReferenceAnswer = &refAnswer
	}

	return updatedQuestion
}

func QuestionToGRPC(question *model.Question) *desc.GetResponse {
	options := make([]*desc.QuestionOption, len(question.Options))
	for i, option := range question.Options {
		options[i] = &desc.QuestionOption{
			Id:        int64(option.Id),
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		}
	}

	res := &desc.GetResponse{
		Id:           int64(question.Id),
		Text:         question.Text,
		QuestionType: desc.QuestionType(question.Type),
		Options:      options,
		CreatedAt:    timestamppb.New(question.CreatedAt),
		UpdatedAt:    timestamppb.New(*question.UpdatedAt),
	}

	if question.ReferenceAnswer != nil {
		res.ReferenceAnswer = wrapperspb.String(*question.ReferenceAnswer)
	}

	if question.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*question.UpdatedAt)
	}

	return res
}

func NewQuestionOptionsFromGRPC(req *desc.AddOptionsRequest) []*model.NewQuestionOption {
	newOptions := make([]*model.NewQuestionOption, len(req.Options))
	for i, option := range req.Options {
		newOptions[i] = &model.NewQuestionOption{
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		}
	}

	return newOptions
}

func DeleteQuestionOptionsFromGRPC(req *desc.DeleteOptionsRequest) []int {
	optionIds := make([]int, len(req.Ids))
	for i, id := range req.Ids {
		optionIds[i] = int(id)
	}

	return optionIds
}

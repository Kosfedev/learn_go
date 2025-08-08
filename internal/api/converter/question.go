package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func QuestionToGRPC(question *model.Question) *desc.GetResponse {
	questionGRPC := &desc.Question{
		Id:           question.ID,
		Text:         question.Text,
		QuestionType: desc.QuestionType(question.Type),
		CreatedAt:    timestamppb.New(question.CreatedAt),
	}

	if question.ReferenceAnswer != nil {
		questionGRPC.ReferenceAnswer = wrapperspb.String(*question.ReferenceAnswer)
	}

	if question.UpdatedAt != nil {
		questionGRPC.UpdatedAt = timestamppb.New(*question.UpdatedAt)
	}

	res := &desc.GetResponse{
		Data: &desc.GetResponse_Data{Question: questionGRPC},
	}

	return res
}

func NewQuestionFromGRPC(req *desc.CreateRequest) *model.NewQuestion {
	newOptions := make([]*model.NewQuestionOption, len(req.Options))
	for i, optionGRPC := range req.Options {
		optionServ := &model.NewQuestionOption{
			Text:      optionGRPC.Text,
			IsCorrect: optionGRPC.IsCorrect,
		}

		newOptions[i] = optionServ
	}

	newQuestion := &model.NewQuestion{
		Text:    req.Text,
		Type:    model.QuestionType(req.QuestionType),
		Options: newOptions,
	}

	if req.ReferenceAnswer != nil {
		refAnswer := req.ReferenceAnswer.GetValue()
		newQuestion.ReferenceAnswer = &refAnswer
	}

	return newQuestion
}

func UpdatedQuestionFromGRPC(req *desc.UpdateRequest) *model.UpdatedQuestion {
	updatedQuestion := &model.UpdatedQuestion{}

	if req.Text != nil {
		text := req.GetText().GetValue()
		updatedQuestion.Text = &text
	}

	if req.QuestionType != nil {
		questionType := model.QuestionType(*req.QuestionType.Enum())
		updatedQuestion.Type = &questionType
	}

	if req.ReferenceAnswer != nil {
		refAnswer := req.ReferenceAnswer.GetValue()
		updatedQuestion.ReferenceAnswer = &refAnswer
	}

	return updatedQuestion
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

package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func QuestionToGRPC(question *model.Question) *desc.GetResponse {
	options := make([]*desc.QuestionOption, len(question.Options))
	for i, option := range question.Options {
		options[i] = &desc.QuestionOption{
			Id:        option.ID,
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		}
	}

	// TODO: унифицировать с конвертером subcategory
	subcategories := make([]*desc.Subcategory, len(question.Subcategories))
	for i, subcategory := range question.Subcategories {
		subcategories[i] = &desc.Subcategory{
			Id:         subcategory.ID,
			Name:       subcategory.Name,
			CategoryId: subcategory.CategoryID,
			CreatedAt:  timestamppb.New(subcategory.CreatedAt),
		}

		if subcategory.UpdatedAt != nil {
			subcategories[i].UpdatedAt = timestamppb.New(*subcategory.UpdatedAt)
		}
	}

	// TODO: унифицировать с конвертером set
	sets := make([]*desc.Set, len(question.Sets))
	for i, set := range question.Sets {
		sets[i] = &desc.Set{
			Id:        set.ID,
			Name:      set.Name,
			CreatedAt: timestamppb.New(set.CreatedAt),
		}

		if set.UpdatedAt != nil {
			subcategories[i].UpdatedAt = timestamppb.New(*set.UpdatedAt)
		}
	}

	res := &desc.GetResponse{
		Id:            question.ID,
		Text:          question.Text,
		QuestionType:  desc.QuestionType(question.Type),
		Options:       options,
		Subcategories: subcategories,
		Sets:          sets,
		CreatedAt:     timestamppb.New(question.CreatedAt),
	}

	if question.ReferenceAnswer != nil {
		res.ReferenceAnswer = wrapperspb.String(*question.ReferenceAnswer)
	}

	if question.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*question.UpdatedAt)
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

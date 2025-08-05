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
			Id:        int64(option.ID),
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		}
	}

	subcategories := make([]*desc.Subcategory, len(question.Subcategories))
	for i, subcategory := range question.Subcategories {
		subcategories[i] = &desc.Subcategory{
			Id:         int64(subcategory.ID),
			Name:       subcategory.Name,
			CategoryId: int64(subcategory.CategoryID),
			CreatedAt:  timestamppb.New(subcategory.CreatedAt),
			UpdatedAt:  nil,
		}

		if subcategory.UpdatedAt != nil {
			subcategories[i].UpdatedAt = timestamppb.New(*subcategory.UpdatedAt)
		}
	}

	res := &desc.GetResponse{
		Id:            int64(question.ID),
		Text:          question.Text,
		QuestionType:  desc.QuestionType(question.Type),
		Options:       options,
		Subcategories: subcategories,
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

func DeleteQuestionOptionsFromGRPC(req *desc.DeleteOptionsRequest) []int {
	optionIds := make([]int, len(req.Ids))
	for i, id := range req.Ids {
		optionIds[i] = int(id)
	}

	return optionIds
}

// TODO: refactor to SubcategoryIdsFromGRPC???
func AddSubcategoriesFromGRPC(req *desc.AddSubcategoriesRequest) []int {
	subcategoryIds := make([]int, len(req.SubcategoryIds))
	for i, id := range req.SubcategoryIds {
		subcategoryIds[i] = int(id)
	}

	return subcategoryIds
}

func RemoveSubcategoriesFromGRPC(req *desc.RemoveSubcategoriesRequest) []int {
	subcategoryIds := make([]int, len(req.SubcategoryIds))
	for i, id := range req.SubcategoryIds {
		subcategoryIds[i] = int(id)
	}

	return subcategoryIds
}

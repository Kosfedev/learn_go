package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_v1"
)

// TODO: переработать к хуям
func QuestionFormToGRPC(questionForm *model.QuestionForm) *desc.GetResponse {
	res := &desc.GetResponse{
		Data: &desc.GetResponse_Data{
			Question:      QuestionToGRPCTemp(questionForm.Question),
			Options:       QuestionOptionsToGRPCTemp(questionForm.Options),
			Sets:          SetsToGRPC(questionForm.Sets),
			Subcategories: SubcategoriesToGRPC(questionForm.Subcategories),
		},
	}

	return res
}

func QuestionToGRPCTemp(question *model.Question) *desc.Question {
	res := &desc.Question{
		Id:           question.ID,
		Text:         question.Text,
		QuestionType: desc.QuestionType(question.Type),
		CreatedAt:    timestamppb.New(question.CreatedAt),
	}

	if question.ReferenceAnswer != nil {
		res.ReferenceAnswer = wrapperspb.String(*question.ReferenceAnswer)
	}

	if question.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*question.UpdatedAt)
	}

	return res
}

func QuestionOptionsToGRPCTemp(options []*model.QuestionOption) []*desc.QuestionOption {
	optionsGRPC := make([]*desc.QuestionOption, len(options))
	for i, option := range options {
		optionsGRPC[i] = &desc.QuestionOption{
			Id:        option.ID,
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		}
	}

	return optionsGRPC
}

func SetsToGRPC(sets []*model.Set) []*desc.Set {
	setsGRPC := make([]*desc.Set, len(sets))
	for i, set := range sets {
		setGRPC := &desc.Set{
			Id:        set.ID,
			Name:      set.Name,
			CreatedAt: timestamppb.New(set.CreatedAt),
		}

		if set.UpdatedAt != nil {
			setGRPC.UpdatedAt = timestamppb.New(*set.UpdatedAt)
		}

		setsGRPC[i] = setGRPC
	}

	return setsGRPC
}

func SubcategoriesToGRPC(subcategories []*model.Subcategory) []*desc.Subcategory {
	subcategoriesGRPC := make([]*desc.Subcategory, len(subcategories))
	for i, subcategory := range subcategories {
		subcategoryGRPC := &desc.Subcategory{
			Id:         subcategory.ID,
			Name:       subcategory.Name,
			CategoryId: subcategory.CategoryID,
			CreatedAt:  timestamppb.New(subcategory.CreatedAt),
			UpdatedAt:  nil,
		}

		if subcategory.UpdatedAt != nil {
			subcategoryGRPC.UpdatedAt = timestamppb.New(*subcategory.UpdatedAt)
		}

		subcategoriesGRPC[i] = subcategoryGRPC
	}

	return subcategoriesGRPC
}

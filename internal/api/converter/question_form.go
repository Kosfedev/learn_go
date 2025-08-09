package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Kosfedev/learn_go/internal/model"
	descQuestionForm "github.com/Kosfedev/learn_go/pkg/question_form_v1"
	descQuestion "github.com/Kosfedev/learn_go/pkg/question_v1"
	descSet "github.com/Kosfedev/learn_go/pkg/set_v1"
	descSubcategory "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

// TODO: переработать к хуям
func QuestionFormToGRPC(questionForm *model.QuestionForm) *descQuestionForm.GetResponse {
	res := &descQuestionForm.GetResponse{
		Data: &descQuestionForm.GetResponse_Data{
			Question:      QuestionToGRPCTemp(questionForm.Question),
			Options:       QuestionOptionsToGRPCTemp(questionForm.Options),
			Sets:          SetsToGRPC(questionForm.Sets),
			Subcategories: SubcategoriesToGRPC(questionForm.Subcategories),
		},
	}

	return res
}

func QuestionToGRPCTemp(question *model.Question) *descQuestion.Question {
	res := &descQuestion.Question{
		Id:           question.ID,
		Text:         question.Text,
		QuestionType: descQuestion.QuestionType(question.Type),
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

func QuestionOptionsToGRPCTemp(options []*model.QuestionOption) []*descQuestionForm.QuestionOption {
	optionsGRPC := make([]*descQuestionForm.QuestionOption, len(options))
	for i, option := range options {
		optionsGRPC[i] = &descQuestionForm.QuestionOption{
			Id:        option.ID,
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		}
	}

	return optionsGRPC
}

func SetsToGRPC(sets []*model.Set) []*descSet.Set {
	setsGRPC := make([]*descSet.Set, len(sets))
	for i, set := range sets {
		setGRPC := &descSet.Set{
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

func SubcategoriesToGRPC(subcategories []*model.Subcategory) []*descSubcategory.Subcategory {
	subcategoriesGRPC := make([]*descSubcategory.Subcategory, len(subcategories))
	for i, subcategory := range subcategories {
		subcategoryGRPC := &descSubcategory.Subcategory{
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

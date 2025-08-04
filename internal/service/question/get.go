package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Get(ctx context.Context, questionId int) (*model.Question, error) {
	question, err := qs.questionRepo.Get(ctx, questionId)
	if err != nil {
		return nil, err
	}

	questionSubcategoryIds, err := qs.questionRepo.ListSubcategoriesByQuestionId(ctx, questionId)
	if err != nil {
		return nil, err
	}

	subcategories, err := qs.subcategoryRepo.ListByIds(ctx, questionSubcategoryIds)
	if err != nil {
		return nil, err
	}

	question.Subcategories = subcategories

	return question, nil
}

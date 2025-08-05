package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Get(ctx context.Context, questionID int) (*model.Question, error) {
	question, err := qs.questionRepo.Get(ctx, questionID)
	if err != nil {
		return nil, err
	}

	questionSubcategoryIDs, err := qs.questionSubcategoryRepo.ListSubcategoriesByQuestionID(ctx, questionID)
	if err != nil {
		return nil, err
	}

	subcategories, err := qs.subcategoryRepo.ListByIDs(ctx, questionSubcategoryIDs)
	if err != nil {
		return nil, err
	}

	question.Subcategories = subcategories

	return question, nil
}

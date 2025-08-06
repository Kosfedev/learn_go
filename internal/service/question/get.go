package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Get(ctx context.Context, questionID int64) (*model.Question, error) {
	question, err := qs.questionRepo.Get(ctx, questionID)
	if err != nil {
		return nil, err
	}

	subcategoryIDs, err := qs.questionSubcategoryRepo.ListSubcategoriesByQuestionID(ctx, questionID)
	if err != nil {
		return nil, err
	}

	subcategories, err := qs.subcategoryRepo.ListByIDs(ctx, subcategoryIDs)
	if err != nil {
		return nil, err
	}

	question.Subcategories = subcategories

	setIDs, err := qs.questionSetRepo.ListSetsByQuestionID(ctx, questionID)
	if err != nil {
		return nil, err
	}

	sets, err := qs.setRepo.ListByIDs(ctx, setIDs)
	if err != nil {
		return nil, err
	}

	question.Sets = sets

	return question, nil
}

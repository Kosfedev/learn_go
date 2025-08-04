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

	// TODO: replace with ListByIds
	subcategory, err := qs.subcategoryRepo.Get(ctx, questionId)
	if err != nil {
		return nil, err
	}

	question.Subcategories = []*model.Subcategory{subcategory}

	return question, nil
}

package question

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) AddSubcategories(ctx context.Context, req *desc.AddSubcategoriesRequest) (*desc.AddSubcategoriesResponse, error) {
	err := questionImpl.questionService.AddSubcategories(ctx, req.QuestionId, req.SubcategoryIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

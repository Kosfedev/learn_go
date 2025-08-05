package question

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) RemoveSubcategories(ctx context.Context, req *desc.RemoveSubcategoriesRequest) (*desc.RemoveSubcategoriesResponse, error) {
	err := questionImpl.questionService.RemoveSubcategories(ctx, req.QuestionId, req.SubcategoryIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

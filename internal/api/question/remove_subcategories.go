package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) RemoveSubcategories(ctx context.Context, req *desc.RemoveSubcategoriesRequest) (*desc.RemoveSubcategoriesResponse, error) {
	subcategoryIds := converter.RemoveSubcategoriesFromGRPC(req)
	err := questionImpl.questionService.RemoveSubcategories(ctx, int(req.QuestionId), subcategoryIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

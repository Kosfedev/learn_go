package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) RemoveSubcategories(ctx context.Context, req *desc.RemoveSubcategoriesRequest) (*desc.RemoveSubcategoriesResponse, error) {
	subcategoryIds := converter.RemoveSubcategoriesFromGRPC(req)
	err := questionImpl.questionService.RemoveSubcategories(ctx, req.QuestionId, subcategoryIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) AddSubcategories(ctx context.Context, req *desc.AddSubcategoriesRequest) (*desc.AddSubcategoriesResponse, error) {
	subcategoryIds := converter.AddSubcategoriesFromGRPC(req)
	err := questionImpl.questionService.AddSubcategories(ctx, int(req.QuestionId), subcategoryIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

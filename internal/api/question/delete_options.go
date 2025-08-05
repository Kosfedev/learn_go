package question

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) DeleteOptions(ctx context.Context, req *desc.DeleteOptionsRequest) (*desc.DeleteOptionsResponse, error) {
	err := questionImpl.questionService.DeleteOptions(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

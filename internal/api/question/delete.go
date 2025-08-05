package question

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	err := questionImpl.questionService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

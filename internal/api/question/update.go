package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	question := converter.UpdatedQuestionFromGRPC(req)
	err := questionImpl.questionService.Update(ctx, int(req.Id), question)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

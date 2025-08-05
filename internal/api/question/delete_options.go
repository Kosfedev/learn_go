package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) DeleteOptions(ctx context.Context, req *desc.DeleteOptionsRequest) (*desc.DeleteOptionsResponse, error) {
	optionIds := converter.DeleteQuestionOptionsFromGRPC(req)
	err := questionImpl.questionService.DeleteOptions(ctx, optionIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

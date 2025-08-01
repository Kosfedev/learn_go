package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	question, err := questionImpl.questionService.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	res := converter.QuestionToGRPC(question)

	return res, nil
}

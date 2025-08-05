package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) AddOptions(ctx context.Context, req *desc.AddOptionsRequest) (*desc.AddOptionsResponse, error) {
	newOptions := converter.NewQuestionOptionsFromGRPC(req)
	err := questionImpl.questionService.AddOptions(ctx, int64(req.QuestionId), newOptions)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

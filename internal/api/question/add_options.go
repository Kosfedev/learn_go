package question

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) AddOptions(ctx context.Context, req *desc.AddOptionsRequest) (*desc.AddOptionsResponse, error) {
	newOptions := converter.NewQuestionOptionsFromGRPC(req)
	log.Printf("question id: %d; new options: %+v", req.QuestionId, newOptions)
	err := questionImpl.questionService.AddOptions(ctx, int(req.QuestionId), newOptions)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

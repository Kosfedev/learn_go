package question

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

func (questionImpl *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	question := converter.NewQuestionFromGRPC(req)
	log.Printf("new question: %+v", question)
	id, err := questionImpl.questionService.Create(ctx, question)
	if err != nil {
		return nil, err
	}

	res := &desc.CreateResponse{
		Id: int64(id),
	}

	return res, nil
}

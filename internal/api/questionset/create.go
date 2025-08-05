package questionset

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/questionset_v1"
)

func (questionSetImpl *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	questionSet := converter.NewQuestionSetFromGRPC(req)
	id, err := questionSetImpl.questionSetService.Create(ctx, questionSet)
	if err != nil {
		return nil, err
	}

	res := &desc.CreateResponse{
		Id: id,
	}

	return res, nil
}

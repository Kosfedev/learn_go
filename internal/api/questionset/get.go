package questionset

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/questionset_v1"
)

func (questionSetImpl *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	questionSet, err := questionSetImpl.questionSetService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res := converter.QuestionSetToGRPC(questionSet)

	return res, nil
}

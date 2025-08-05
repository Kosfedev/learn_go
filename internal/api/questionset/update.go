package questionset

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/questionset_v1"
)

func (questionSetImpl *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	questionSet := converter.UpdatedQuestionSetFromGRPC(req)
	err := questionSetImpl.questionSetService.Update(ctx, int(req.Id), questionSet)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

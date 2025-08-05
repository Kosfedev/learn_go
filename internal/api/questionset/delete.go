package questionset

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/questionset_v1"
)

func (questionSetImpl *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	err := questionSetImpl.questionSetService.Delete(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

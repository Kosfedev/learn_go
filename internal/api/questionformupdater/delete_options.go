package questionformupdater

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func (impl *Implementation) DeleteOptions(ctx context.Context, req *desc.DeleteOptionsRequest) (*desc.DeleteOptionsResponse, error) {
	err := impl.questionFormUpdaterService.DeleteOptions(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

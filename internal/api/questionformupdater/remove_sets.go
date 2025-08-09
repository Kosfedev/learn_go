package questionformupdater

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func (impl *Implementation) RemoveSets(ctx context.Context, req *desc.RemoveSetsRequest) (*desc.RemoveSetsResponse, error) {
	err := impl.questionFormUpdaterService.RemoveSets(ctx, req.QuestionId, req.SetIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

package questionformupdater

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func (impl *Implementation) AddSets(ctx context.Context, req *desc.AddSetsRequest) (*desc.AddSetsResponse, error) {
	err := impl.questionFormUpdaterService.AddSets(ctx, req.QuestionId, req.SetIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

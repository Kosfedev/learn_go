package questionformupdater

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func (impl *Implementation) AddOptions(ctx context.Context, req *desc.CreateOptionsRequest) (*desc.CreateOptionsResponse, error) {
	newOptions := converter.NewQuestionOptionsFromGRPC(req.Options)
	err := impl.questionFormUpdaterService.CreateOptions(ctx, req.QuestionId, newOptions)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

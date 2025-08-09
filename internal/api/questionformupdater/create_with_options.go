package questionformupdater

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func (i *Implementation) CreateWithOptions(ctx context.Context, req *desc.CreateWithOptionsRequest) (*desc.CreateWithOptionsResponse, error) {
	newQuestionWithOptions := converter.NewQuestionWithOptionsFromGRPC(req)
	id, err := i.questionFormUpdaterService.CreateWithOptions(ctx, newQuestionWithOptions)
	if err != nil {
		return nil, err
	}

	res := &desc.CreateWithOptionsResponse{
		Id: id,
	}

	return res, err
}

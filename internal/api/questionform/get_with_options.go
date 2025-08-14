package questionform

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_v1"
)

func (impl *Implementation) GetWithOptions(ctx context.Context, req *desc.GetWithOptionsRequest) (*desc.GetWithOptionsResponse, error) {
	questionWithOptions, err := impl.questionFormService.GetWithOptions(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	questionWithOptionsGRPC := converter.QuestionWithOptionsToGRPC(questionWithOptions)

	return questionWithOptionsGRPC, err
}

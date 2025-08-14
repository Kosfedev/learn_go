package questionform

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_v1"
)

func (impl *Implementation) GetWithOptionsSetsSubcategories(ctx context.Context, req *desc.GetFormRequest) (*desc.GetFormResponse, error) {
	questionForm, err := impl.questionFormService.GetWithOptionsSetsSubcategories(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	questionFormGRPC := converter.QuestionFormToGRPC(questionForm)

	return questionFormGRPC, err
}

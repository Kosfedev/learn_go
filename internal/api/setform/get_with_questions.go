package setform

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/set_form_v1"
)

func (impl *Implementation) GetWithQuestions(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	setForm, err := impl.setFormService.GetWithQuestions(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	setFormGRPC := converter.SetFormToGRPC(setForm)

	return setFormGRPC, err
}

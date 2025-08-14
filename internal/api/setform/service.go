package setform

import (
	desc "github.com/Kosfedev/learn_go/pkg/set_form_v1"

	"github.com/Kosfedev/learn_go/internal/service"
)

type Implementation struct {
	desc.UnimplementedSetFormV1Server
	setFormService service.SetFormService
}

func NewImplementation(questionFormService service.SetFormService) *Implementation {
	return &Implementation{
		setFormService: questionFormService,
	}
}

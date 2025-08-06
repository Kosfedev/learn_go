package questionset

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/questionset_v1"
)

type Implementation struct {
	desc.UnimplementedQuestionSetV1Server
	questionSetService service.SetService
}

func NewImplementation(questionSetService service.SetService) *Implementation {
	return &Implementation{
		questionSetService: questionSetService,
	}
}

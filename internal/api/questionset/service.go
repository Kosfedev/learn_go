package questionset

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/questionset_v1"
)

type Implementation struct {
	desc.UnimplementedQuestionSetV1Server
	questionSetService service.QuestionSetService
}

func NewImplementation(questionSetService service.QuestionSetService) *Implementation {
	return &Implementation{
		questionSetService: questionSetService,
	}
}

package questionform

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_v1"
)

type Implementation struct {
	desc.UnimplementedQuestionFormV1Server
	questionFormService service.QuestionFormService
}

func NewImplementation(questionFormService service.QuestionFormService) *Implementation {
	return &Implementation{
		questionFormService: questionFormService,
	}
}

package question

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

type Implementation struct {
	desc.UnimplementedQuestionV1Server
	questionService service.QuestionService
}

func NewImplementation(questionService service.QuestionService) *Implementation {
	return &Implementation{
		questionService: questionService,
	}
}

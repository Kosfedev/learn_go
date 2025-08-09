package questionformupdater

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

type Implementation struct {
	desc.UnimplementedQuestionFormUpdaterV1Server
	questionFormUpdaterService service.QuestionFormUpdaterService
}

func NewImplementation(questionFormUpdaterService service.QuestionFormUpdaterService) *Implementation {
	return &Implementation{
		questionFormUpdaterService: questionFormUpdaterService,
	}
}

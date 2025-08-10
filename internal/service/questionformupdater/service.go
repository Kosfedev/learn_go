package questionformupdater

import (
	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	txManager               db.TxManager
	questionRepo            repository.QuestionRepository
	questionOptionRepo      repository.QuestionOptionRepository
	questionSubcategoryRepo repository.QuestionSubcategoryRepository
	questionSetRepo         repository.QuestionSetForQuestionRepository
}

func NewService(
	txManager db.TxManager,
	questionRepo repository.QuestionRepository,
	questionOptionRepo repository.QuestionOptionRepository,
	questionSubcategoryRepo repository.QuestionSubcategoryRepository,
	questionSetRepo repository.QuestionSetForQuestionRepository,
) service.QuestionFormUpdaterService {
	return &serv{
		txManager:               txManager,
		questionRepo:            questionRepo,
		questionOptionRepo:      questionOptionRepo,
		questionSubcategoryRepo: questionSubcategoryRepo,
		questionSetRepo:         questionSetRepo,
	}
}

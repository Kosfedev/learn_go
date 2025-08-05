package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
)

func (r *repo) AddOptions(ctx context.Context, questionID int, options []*model.NewQuestionOption) error {
	if len(options) == 0 {
		return nil
	}

	var values []interface{}
	nArgs := 3
	queryRaw := `INSERT INTO question_option (question_id, text, is_correct) VALUES`

	for i, option := range options {
		queryRaw += fmt.Sprintf(" ($%d, $%d, $%d),", i*nArgs+1, i*nArgs+2, i*nArgs+3)
		newOptionRepo := converter.NewQuestionOptionToPGSQL(questionID, option)
		values = append(values, newOptionRepo.QuestionID, newOptionRepo.Text, newOptionRepo.IsCorrect)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	query := db.Query{
		Name:     "question_repository.add_options",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/questionset/pg/converter"
)

func (r *repo) Create(ctx context.Context, newQuestionSet *model.NewQuestionSet) (int, error) {
	newQuestionSetRepo := converter.NewQuestionSetToPGSQL(newQuestionSet)
	query := `
		INSERT INTO question_set (name) 
		VALUES ($1)
		RETURNING id;`

	var questionSetIdRepo int32
	err := r.db.QueryRowContext(ctx, query, newQuestionSetRepo.Name).Scan(&questionSetIdRepo)
	if err != nil {
		return 0, err
	}

	questionSetId := int(questionSetIdRepo)
	// TODO: вынести в общую функцию?
	// TODO: нужна транзакция
	if len(newQuestionSet.QuestionIds) > 0 {
		var values []interface{}
		nArgs := 2
		query = `INSERT INTO question_set_question (question_id, question_set_id) VALUES`

		for i, questionId := range newQuestionSet.QuestionIds {
			query += fmt.Sprintf(" ($%d, $%d),", i*nArgs+1, i*nArgs+2)
			questionIdRepo := int32(questionId)
			values = append(values, questionIdRepo, questionSetIdRepo)
		}

		query = strings.TrimSuffix(query, ",")

		_, err = r.db.Query(query, values...)
		if err != nil {
			return 0, err
		}
	}

	return questionSetId, nil
}

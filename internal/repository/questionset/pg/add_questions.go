package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) addQuestions(ctx context.Context, questionSetId int, questionIds []int) error {
	if len(questionIds) == 0 {
		return nil
	}

	var values []interface{}
	nArgs := 2
	questionSetIdRepo := int32(questionSetId)
	queryRaw := `INSERT INTO question_set_question (question_id, question_set_id) VALUES`

	for i, questionId := range questionIds {
		queryRaw += fmt.Sprintf(" ($%d, $%d),", i*nArgs+1, i*nArgs+2)
		questionIdRepo := int32(questionId)
		values = append(values, questionIdRepo, questionSetIdRepo)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	query := db.Query{
		Name:     "question_set_repository.add_questions",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) addQuestions(ctx context.Context, questionSetID int, questionIDs []int) error {
	if len(questionIDs) == 0 {
		return nil
	}

	var values []interface{}
	nArgs := 2
	questionSetIDRepo := int32(questionSetID)
	queryRaw := `INSERT INTO question_set_question (question_id, question_set_id) VALUES`

	for i, questionID := range questionIDs {
		queryRaw += fmt.Sprintf(" ($%d, $%d),", i*nArgs+1, i*nArgs+2)
		questionIDRepo := int32(questionID)
		values = append(values, questionIDRepo, questionSetIDRepo)
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

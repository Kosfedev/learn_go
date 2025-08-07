package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) addQuestions(ctx context.Context, setID int64, questionIDs []int64) error {
	if len(questionIDs) == 0 {
		return nil
	}

	var values []interface{}
	nArgs := 2
	questionSetIDRepo := setID
	queryRaw := `INSERT INTO question_set (question_id, set_id) VALUES`

	for i, questionID := range questionIDs {
		queryRaw += fmt.Sprintf(" ($%d, $%d),", i*nArgs+1, i*nArgs+2)
		questionIDRepo := questionID
		values = append(values, questionIDRepo, questionSetIDRepo)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	query := db.Query{
		Name:     "set_repository.add_questions",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

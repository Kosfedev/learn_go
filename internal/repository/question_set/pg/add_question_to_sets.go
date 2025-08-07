package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) AddQuestionToSets(ctx context.Context, questionID int64, setIDs []int64) error {
	if len(setIDs) == 0 {
		return nil
	}

	var values []interface{}
	nArgs := 2
	queryRaw := `INSERT INTO question_set (question_id, set_id) VALUES`

	for i, setID := range setIDs {
		queryRaw += fmt.Sprintf(" ($%d, $%d),", i*nArgs+1, i*nArgs+2)
		values = append(values, setID, questionID)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	query := db.Query{
		Name:     "question_set_repository.add_question_to_sets",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

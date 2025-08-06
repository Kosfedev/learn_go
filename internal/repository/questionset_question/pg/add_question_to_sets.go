package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) AddQuestionToSets(ctx context.Context, questionID int64, questionSetIDs []int64) error {
	if len(questionSetIDs) == 0 {
		return nil
	}

	var values []interface{}
	nArgs := 2
	queryRaw := `INSERT INTO question_subcategory (question_id, subcategory_id) VALUES`

	for i, questionSetID := range questionSetIDs {
		queryRaw += fmt.Sprintf(" ($%d, $%d),", i*nArgs+1, i*nArgs+2)
		values = append(values, questionSetID, questionID)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	query := db.Query{
		Name:     "question_question_set_repository.add_question_to_sets",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

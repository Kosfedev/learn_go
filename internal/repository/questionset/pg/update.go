package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/questionset/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int64, updatedQuestionSet *model.UpdatedQuestionSet) error {
	updatedQuestionSetRepo := converter.UpdatedQuestionSetToPGSQL(updatedQuestionSet)
	index := 2
	values := []interface{}{id}
	queryRaw := "UPDATE question_set SET"

	if updatedQuestionSetRepo.Name.Valid {
		values = append(values, updatedQuestionSetRepo.Name)
		queryRaw += fmt.Sprintf(" name = $%d,", index)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	queryRaw += " WHERE id = $1"
	query := db.Query{
		Name:     "question_set_repository.update",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

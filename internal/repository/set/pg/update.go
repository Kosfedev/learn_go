package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/set/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int64, updatedSet *model.UpdatedSet) error {
	updatedSetRepo := converter.UpdatedSetToPGSQL(updatedSet)
	index := 2
	values := []interface{}{id}
	queryRaw := "UPDATE set SET"

	if updatedSetRepo.Name.Valid {
		values = append(values, updatedSetRepo.Name)
		queryRaw += fmt.Sprintf(" name = $%d,", index)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	queryRaw += " WHERE id = $1"
	query := db.Query{
		Name:     "set_repository.update",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

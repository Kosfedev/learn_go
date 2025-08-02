package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/domain/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, domain *model.UpdatedDomain) error {
	domainRepo := converter.UpdatedDomainToPGSQL(domain)
	values := []interface{}{id}
	query := "UPDATE domain SET"

	if domainRepo.Name.Valid {
		values = append(values, domainRepo.Name)
		query += fmt.Sprintf(" name = $%d,", 2)
	}

	query = strings.TrimSuffix(query, ",")
	query += " WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

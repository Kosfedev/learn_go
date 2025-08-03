package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/domain/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, domain *model.UpdatedDomain) error {
	domainRepo := converter.UpdatedDomainToPGSQL(domain)
	values := []interface{}{id}
	queryRaw := "UPDATE domain SET"

	if domainRepo.Name.Valid {
		values = append(values, domainRepo.Name)
		queryRaw += fmt.Sprintf(" name = $%d,", 2)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	queryRaw += " WHERE id = $1"
	query := db.Query{
		Name:     "domain_repository.update",
		QueryRaw: queryRaw,
	}
	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

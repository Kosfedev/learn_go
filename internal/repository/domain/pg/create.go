package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/domain/pg/converter"
)

func (r *repo) Create(ctx context.Context, domain *model.NewDomain) (int, error) {
	domainRepo := converter.NewDomainToPGSQL(domain)
	query := `INSERT INTO domain(name) VALUES ($1) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, domainRepo.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

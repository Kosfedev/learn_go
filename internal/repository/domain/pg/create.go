package pg

import (
	"context"
	"fmt"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/domain/pg/converter"
)

func (r *repo) Create(ctx context.Context, domain *model.NewDomain) (int, error) {
	domainRepo := converter.NewDomainToPGSQL(domain)
	query := db.Query{
		Name:     "domain_repository.create",
		QueryRaw: `INSERT INTO domain(name) VALUES ($1) RETURNING id`,
	}

	var id int
	fmt.Printf("db: %+v\n", r.db)
	err := r.db.DB().ScanOne(ctx, &id, query, domainRepo.Name)
	if err != nil {
		return 0, err
	}

	return id, nil
}

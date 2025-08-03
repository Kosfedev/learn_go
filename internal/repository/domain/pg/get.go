package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/domain/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/domain/pg/model"
)

func (r *repo) Get(ctx context.Context, id int) (*model.Domain, error) {
	query := db.Query{
		Name:     "domain_get",
		QueryRaw: `SELECT * FROM domain WHERE id = $1`,
	}

	domainRepo := &modelRepo.Domain{}
	err := r.db.DB().ScanOne(ctx, domainRepo, query, id)
	if err != nil {
		return nil, err
	}

	domain := converter.DomainFromPGSQL(domainRepo)

	return domain, nil
}

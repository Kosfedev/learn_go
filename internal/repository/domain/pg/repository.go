package pg

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/domain/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/domain/pg/model"
)

const (
	tableDomain = "domain"
	columnID    = "id"
	columnName  = "name"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.DomainRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, domain *model.NewDomain) (int64, error) {
	domainRepo := converter.NewDomainToPGSQL(domain)
	builderInsert := sq.Insert(tableDomain).
		PlaceholderFormat(sq.Dollar).
		Columns(columnName).
		Values(domainRepo.Name).
		Suffix(fmt.Sprintf("RETURNING %v", columnID))

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	query := db.Query{
		Name:     "domain_repository.create",
		QueryRaw: queryRaw,
	}

	var id int64
	fmt.Printf("db: %+v\n", r.db)
	err = r.db.DB().ScanOne(ctx, &id, query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Domain, error) {
	builderSelect := sq.Select(columnName).
		From(tableDomain).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "domain_repository.get",
		QueryRaw: queryRaw,
	}

	domainRepo := &modelRepo.Domain{}
	err = r.db.DB().ScanOne(ctx, domainRepo, query, args...)
	if err != nil {
		return nil, err
	}

	domain := converter.DomainFromPGSQL(domainRepo)

	return domain, nil
}

func (r *repo) Update(ctx context.Context, id int64, domain *model.UpdatedDomain) error {
	domainRepo := converter.UpdatedDomainToPGSQL(domain)
	builderUpdate := sq.Update(tableDomain).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	if domainRepo.Name.Valid {
		builderUpdate = builderUpdate.Set(columnName, domainRepo.Name)
	}

	queryRaw, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "domain_repository.update",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableDomain).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "domain_repository.delete",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

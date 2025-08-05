package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/domain/pg/model"
)

func DomainFromPGSQL(domainRepo *modelRepo.Domain) *model.Domain {
	domain := &model.Domain{
		ID:        domainRepo.ID,
		Name:      domainRepo.Name,
		CreatedAt: domainRepo.CreatedAt,
	}

	if domainRepo.UpdatedAt.Valid {
		domain.UpdatedAt = &domainRepo.UpdatedAt.Time
	}

	return domain
}

func NewDomainToPGSQL(domain *model.NewDomain) *modelRepo.NewDomain {
	domainRepo := &modelRepo.NewDomain{
		Name: domain.Name,
	}

	return domainRepo
}

func UpdatedDomainToPGSQL(domain *model.UpdatedDomain) *modelRepo.UpdatedDomain {
	domainRepo := &modelRepo.UpdatedDomain{}

	if domain.Name != nil {
		domainRepo.Name = sql.NullString{
			String: *domain.Name,
			Valid:  true,
		}
	}

	return domainRepo
}

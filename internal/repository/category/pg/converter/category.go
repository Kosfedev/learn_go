package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/category/pg/model"
)

func CategoryFromPGSQL(categoryRepo *modelRepo.Category) *model.Category {
	category := &model.Category{
		Id:        int(categoryRepo.Id),
		Name:      categoryRepo.Name,
		DomainId:  int(categoryRepo.DomainId),
		CreatedAt: categoryRepo.CreatedAt,
	}

	if categoryRepo.UpdatedAt.Valid {
		category.UpdatedAt = &categoryRepo.UpdatedAt.Time
	}

	return category
}

func NewCategoryToPGSQL(category *model.NewCategory) *modelRepo.NewCategory {
	categoryRepo := &modelRepo.NewCategory{
		Name:     category.Name,
		DomainId: int32(category.DomainId),
	}

	return categoryRepo
}

func UpdatedCategoryToPGSQL(category *model.UpdatedCategory) *modelRepo.UpdatedCategory {
	categoryRepo := &modelRepo.UpdatedCategory{}

	if category.Name != nil {
		categoryRepo.Name = sql.NullString{
			String: *category.Name,
			Valid:  true,
		}
	}

	if category.DomainId != nil {
		categoryRepo.DomainId = sql.NullInt32{
			Int32: int32(*category.DomainId),
			Valid: true,
		}
	}

	return categoryRepo
}

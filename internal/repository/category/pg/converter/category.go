package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/category/pg/model"
)

func CategoryFromPGSQL(categoryRepo *modelRepo.Category) *model.Category {
	category := &model.Category{
		ID:        int64(categoryRepo.ID),
		Name:      categoryRepo.Name,
		DomainID:  int64(categoryRepo.DomainID),
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
		DomainID: int32(category.DomainID),
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

	if category.DomainID != nil {
		categoryRepo.DomainID = sql.NullInt32{
			Int32: int32(*category.DomainID),
			Valid: true,
		}
	}

	return categoryRepo
}

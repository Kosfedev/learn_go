package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/model"
)

func SubcategoryFromPGSQL(subcategoryRepo *modelRepo.Subcategory) *model.Subcategory {
	subcategory := &model.Subcategory{
		Id:         int(subcategoryRepo.Id),
		Name:       subcategoryRepo.Name,
		CategoryId: int(subcategoryRepo.CategoryId),
		CreatedAt:  subcategoryRepo.CreatedAt,
	}

	if subcategoryRepo.UpdatedAt.Valid {
		subcategory.UpdatedAt = &subcategoryRepo.UpdatedAt.Time
	}

	return subcategory
}

func NewSubcategoryToPGSQL(subcategory *model.NewSubcategory) *modelRepo.NewSubcategory {
	subcategoryRepo := &modelRepo.NewSubcategory{
		Name:       subcategory.Name,
		CategoryId: int32(subcategory.CategoryId),
	}

	return subcategoryRepo
}

func UpdatedSubcategoryToPGSQL(subcategory *model.UpdatedSubcategory) *modelRepo.UpdatedSubcategory {
	subcategoryRepo := &modelRepo.UpdatedSubcategory{}

	if subcategory.Name != nil {
		subcategoryRepo.Name = sql.NullString{
			String: *subcategory.Name,
			Valid:  true,
		}
	}

	if subcategory.CategoryId != nil {
		subcategoryRepo.CategoryId = sql.NullInt32{
			Int32: int32(*subcategory.CategoryId),
			Valid: true,
		}
	}

	return subcategoryRepo
}

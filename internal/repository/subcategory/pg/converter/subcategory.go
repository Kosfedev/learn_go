package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/model"
)

func SubcategoryFromPGSQL(subcategoryRepo *modelRepo.Subcategory) *model.Subcategory {
	subcategory := &model.Subcategory{
		ID:         int(subcategoryRepo.ID),
		Name:       subcategoryRepo.Name,
		CategoryID: int(subcategoryRepo.CategoryID),
		CreatedAt:  subcategoryRepo.CreatedAt,
	}

	if subcategoryRepo.UpdatedAt.Valid {
		subcategory.UpdatedAt = &subcategoryRepo.UpdatedAt.Time
	}

	return subcategory
}

func SubcategoriesFromPGSQL(subcategoryRepo []*modelRepo.Subcategory) []*model.Subcategory {
	subcategories := make([]*model.Subcategory, len(subcategoryRepo))
	for i, subcategory := range subcategoryRepo {
		subcategories[i] = SubcategoryFromPGSQL(subcategory)
	}

	return subcategories
}

func NewSubcategoryToPGSQL(subcategory *model.NewSubcategory) *modelRepo.NewSubcategory {
	subcategoryRepo := &modelRepo.NewSubcategory{
		Name:       subcategory.Name,
		CategoryID: int32(subcategory.CategoryID),
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

	if subcategory.CategoryID != nil {
		subcategoryRepo.CategoryID = sql.NullInt32{
			Int32: int32(*subcategory.CategoryID),
			Valid: true,
		}
	}

	return subcategoryRepo
}

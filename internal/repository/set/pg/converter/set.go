package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/set/pg/model"
)

func SetFromPGSQL(questionSetRepo *modelRepo.Set) *model.Set {
	questionSetServ := &model.Set{
		ID:        questionSetRepo.ID,
		Name:      questionSetRepo.Name,
		CreatedAt: questionSetRepo.CreatedAt,
	}

	if questionSetRepo.UpdatedAt.Valid {
		questionSetServ.UpdatedAt = &questionSetRepo.UpdatedAt.Time
	}

	return questionSetServ
}

func NewSetToPGSQL(questionSetServ *model.NewSet) *modelRepo.NewSet {
	questionSetRepo := &modelRepo.NewSet{
		Name: questionSetServ.Name,
	}

	return questionSetRepo
}

func UpdatedSetToPGSQL(questionSetServ *model.UpdatedSet) *modelRepo.UpdatedSet {
	questionSetRepo := &modelRepo.UpdatedSet{}

	if questionSetServ.Name != nil {
		questionSetRepo.Name = sql.NullString{
			String: *questionSetServ.Name,
			Valid:  true,
		}
	}

	return questionSetRepo
}

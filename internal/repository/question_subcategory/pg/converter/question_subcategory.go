package converter

import (
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/model"
)

func QuestionSubcategoriesIDsFromPGSQL(questionSubcategoriesRepo []*modelRepo.QuestionSubcategory) []int64 {
	questionSubcategories := make([]int64, len(questionSubcategoriesRepo))

	for i, questionSubcategoryRepo := range questionSubcategoriesRepo {
		questionSubcategories[i] = int64(questionSubcategoryRepo.SubcategoryID)
	}

	return questionSubcategories
}

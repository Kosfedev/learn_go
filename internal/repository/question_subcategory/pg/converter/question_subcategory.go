package converter

import (
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/model"
)

func QuestionSubcategoriesIdsFromPGSQL(questionSubcategoriesRepo []*modelRepo.QuestionSubcategory) []int {
	questionSubcategories := make([]int, len(questionSubcategoriesRepo))

	for i, questionSubcategoryRepo := range questionSubcategoriesRepo {
		questionSubcategories[i] = int(questionSubcategoryRepo.SubcategoryId)
	}

	return questionSubcategories
}

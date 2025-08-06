package converter

import (
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_set/pg/model"
)

func QuestionIDsFromPGSQL(questionQuestionSetsRepo []*modelRepo.QuestionSet) []int64 {
	questionIDs := make([]int64, len(questionQuestionSetsRepo))

	for i, questionQuestionSetRepo := range questionQuestionSetsRepo {
		questionIDs[i] = questionQuestionSetRepo.QuestionID
	}

	return questionIDs
}

func QuestionSetsIDsFromPGSQL(questionQuestionSetsRepo []*modelRepo.QuestionSet) []int64 {
	questionSetIDs := make([]int64, len(questionQuestionSetsRepo))

	for i, questionQuestionSetRepo := range questionQuestionSetsRepo {
		questionSetIDs[i] = questionQuestionSetRepo.SetID
	}

	return questionSetIDs
}

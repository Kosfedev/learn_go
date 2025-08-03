package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/questionset_v1"
)

func QuestionSetToGRPC(question *model.QuestionSet) *desc.GetResponse {
	questionIdsGRPC := make([]int64, len(question.QuestionIds))
	for i, questionId := range question.QuestionIds {
		questionIdsGRPC[i] = int64(questionId)
	}

	res := &desc.GetResponse{
		Id:          int64(question.Id),
		Name:        question.Name,
		QuestionIds: questionIdsGRPC,
		CreatedAt:   timestamppb.New(question.CreatedAt),
	}

	if question.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*question.UpdatedAt)
	}

	return res
}

func NewQuestionSetFromGRPC(req *desc.CreateRequest) *model.NewQuestionSet {
	newQuestionIds := make([]int, len(req.QuestionIds))
	for i, questionIdGRPC := range req.QuestionIds {
		newQuestionIds[i] = int(questionIdGRPC)
	}

	newQuestionSet := &model.NewQuestionSet{
		Name:        req.Name,
		QuestionIds: newQuestionIds,
	}

	return newQuestionSet
}

func UpdatedQuestionSetFromGRPC(req *desc.UpdateRequest) *model.UpdatedQuestionSet {
	updatedQuestionSet := &model.UpdatedQuestionSet{}

	if req.Name != nil {
		name := req.Name.GetValue()
		updatedQuestionSet.Name = &name
	}

	return updatedQuestionSet
}

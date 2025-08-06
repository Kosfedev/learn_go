package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/questionset_v1"
)

func QuestionSetToGRPC(question *model.Set) *desc.GetResponse {
	res := &desc.GetResponse{
		Id:          question.ID,
		Name:        question.Name,
		QuestionIds: question.QuestionIDs,
		CreatedAt:   timestamppb.New(question.CreatedAt),
	}

	if question.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*question.UpdatedAt)
	}

	return res
}

func NewQuestionSetFromGRPC(req *desc.CreateRequest) *model.NewSet {
	newQuestionSet := &model.NewSet{
		Name:        req.Name,
		QuestionIDs: req.QuestionIds,
	}

	return newQuestionSet
}

func UpdatedQuestionSetFromGRPC(req *desc.UpdateRequest) *model.UpdatedSet {
	updatedQuestionSet := &model.UpdatedSet{}

	if req.Name != nil {
		name := req.Name.GetValue()
		updatedQuestionSet.Name = &name
	}

	return updatedQuestionSet
}

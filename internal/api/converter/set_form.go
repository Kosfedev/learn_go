package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Kosfedev/learn_go/internal/model"
	descQuestion "github.com/Kosfedev/learn_go/pkg/question_v1"
	descSetForm "github.com/Kosfedev/learn_go/pkg/set_form_v1"
	descSet "github.com/Kosfedev/learn_go/pkg/set_v1"
)

// TODO: переработать к хуям
func SetFormToGRPC(setForm *model.SetForm) *descSetForm.GetResponse {
	res := &descSetForm.GetResponse{
		Data: &descSetForm.GetResponse_Data{
			Set:       SetToGRPCTemp(setForm.Set),
			Questions: QuestionsToGRPCTemp(setForm.Questions),
		},
	}

	return res
}

func QuestionsToGRPCTemp(questions []*model.Question) []*descQuestion.Question {
	questionsGRPC := make([]*descQuestion.Question, len(questions))
	for i, question := range questions {
		questionGRPC := &descQuestion.Question{
			Id:           question.ID,
			Text:         question.Text,
			QuestionType: descQuestion.QuestionType(question.Type),
			CreatedAt:    timestamppb.New(question.CreatedAt),
		}

		if question.ReferenceAnswer != nil {
			questionGRPC.ReferenceAnswer = wrapperspb.String(*question.ReferenceAnswer)
		}

		if question.UpdatedAt != nil {
			questionGRPC.UpdatedAt = timestamppb.New(*question.UpdatedAt)
		}

		questionsGRPC[i] = questionGRPC
	}

	return questionsGRPC
}

func SetToGRPCTemp(set *model.Set) *descSet.Set {
	setGRPC := &descSet.Set{
		Id:        set.ID,
		Name:      set.Name,
		CreatedAt: timestamppb.New(set.CreatedAt),
	}

	if set.UpdatedAt != nil {
		setGRPC.UpdatedAt = timestamppb.New(*set.UpdatedAt)
	}

	return setGRPC
}

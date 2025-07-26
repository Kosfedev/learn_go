package repository

import "context"

type QuestionRepository interface {
	Create(ctx context.Context)
}

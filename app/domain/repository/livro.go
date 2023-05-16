package repository

import (
	"api/app/domain/model"
	"context"
)

type Livro interface {
	Save(ctx context.Context, l *model.Livro) error
}

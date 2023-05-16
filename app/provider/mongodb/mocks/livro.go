package mocks

import (
	"api/app/domain/model"
	"api/app/domain/repository"
	"context"
)

type Livro struct{}

func NewLivro() repository.Livro {
	return &Livro{}
}

func (a Livro) Save(ctx context.Context, attr *model.Livro) error {
	return nil
}

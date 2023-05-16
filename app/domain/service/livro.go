package service

import (
	"api/app/domain/repository"
	"api/app/server/rest/model/request"
	"context"
)

type LivroService struct {
	repository repository.Livro
}

type LivroServiceInterface interface {
	Save(ctx context.Context, livro request.LivroRequest) error
}

func NewLivroService(repository repository.Livro) LivroServiceInterface {

	return &LivroService{
		repository: repository,
	}

}

func (l *LivroService) Save(ctx context.Context, livro request.LivroRequest) error {
	return nil
}

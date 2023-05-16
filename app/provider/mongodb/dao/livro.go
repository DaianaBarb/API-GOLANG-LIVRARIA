package provider

import (
	"api/app/domain/model"
	"api/app/domain/repository"
	"api/app/provider/mongodb/document"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Livro struct {
	db      *mongo.Database
	options *Options
}

func NewLivroRepository(db *mongo.Database, opt *Options) repository.Livro {
	return &Livro{
		db:      db,
		options: opt,
	}

}

func (l *Livro) Save(ctx context.Context, livro *model.Livro) error {

	doc := document.Livro{
		Nome:      livro.Nome,
		Autor:     livro.Autor,
		Editora:   livro.Editora,
		AnoFabric: livro.AnoFabric,
	}

	_, err := l.db.Collection(l.options.LivroCollection).InsertOne(ctx, &doc)
	if err != nil {

		return err
	}

	return nil

}

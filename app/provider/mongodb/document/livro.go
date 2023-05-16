package document

import "go.mongodb.org/mongo-driver/bson/primitive"

type Livro struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Nome      string              `bson:"name,omitempty"`
	Autor     string              `bson:"autor,omitempty"`
	Editora   string              `bson:"editora,omitempty"`
	AnoFabric string              `bson:"anoFabric,omitempty"`
}

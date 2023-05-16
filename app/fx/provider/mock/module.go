package mock

import (
	"api/app/provider/mongodb/mocks"

	"go.uber.org/fx"
)

func NewLivroModule() fx.Option {
	return fx.Options(
		fx.Provide(
			mocks.NewLivro,
		),
	)
}

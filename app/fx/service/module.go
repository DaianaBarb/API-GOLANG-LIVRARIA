package service

import (
	"api/app/domain/service"

	"go.uber.org/fx"
)

func LivroModuleService() fx.Option {
	return fx.Options(
		fx.Provide(

			service.NewLivroService,
		),
	)
}

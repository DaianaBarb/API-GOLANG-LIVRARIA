package server

import (
	"api/app/domain/service"
	"api/app/fx/provider"
	"api/app/fx/server/rest"

	"github.com/americanas-go/config"
	log "github.com/americanas-go/ignite/americanas-go/log.v1"
	gifx "github.com/americanas-go/ignite/go.uber.org/fx.v1"
	serverfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/americanas-go/multiserver.v1"
	"go.uber.org/fx"
)

func Start() {

	config.Load()
	log.New()

	module := fx.Options(
		provider.LivroModuleRepository(),
		//service.LivroModuleService(),
		fx.Provide(
			service.NewLivroService,
		),
		rest.Module(),
	)

	gifx.NewApp(
		module,

		serverfx.Module(),
	).Run()
}

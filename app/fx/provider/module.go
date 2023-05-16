package provider

import (
	"api/app/domain/repository"
	"api/app/fx/provider/mock"
	"api/app/fx/provider/mongo"

	"go.uber.org/fx"
)

func LivroModuleRepository() fx.Option {
	value := repository.ModuleProviderValue()
	if value == "mongodb" {

		return mongo.NewLivroRepositoryModule()
	}

	return mock.NewLivroModule()

}

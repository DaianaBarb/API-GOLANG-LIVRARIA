package rest

import (
	"api/app/fx/server/rest/routes"
	"api/app/server/rest/handler"

	fiberfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/gofiber/fiber.v2"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	health "github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/contrib/americanas-go/health.v1"

	//	status "github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/contrib/americanas-go/rest-response.v1" //nolint
	//prometheus "github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/contrib/prometheus/client_golang.v1"
	validator "github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		plugins(),
		fiberfx.Module(),
		fx.Provide(validator.New),
		handlers(),
		fx.Invoke(routes.RegisterRouters),
	)
}

func handlers() fx.Option {
	return fx.Options(
		// chamar os handlers
		LivroHandler(),
	)
}

func LivroHandler() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Group:  "handlers",
				Target: handler.NewLivro,
			},
		),
	)
}

func plugins() fx.Option {
	return fx.Options(
		fx.Provide(
			func() []fiber.Plugin {
				return []fiber.Plugin{
					health.Register,
					//	status.Register,
					//prometheus.Register,
					//recover.Register,
				}
			},
		),
	)
}

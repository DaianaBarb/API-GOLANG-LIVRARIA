package routes

import (
	rest "api/app/server/rest"

	fiber "github.com/americanas-go/ignite/gofiber/fiber.v2"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Handlers      []rest.Handler      `group:"handlers"`
	HandlersLivro []rest.HandlerLivro `group:"handlers"`
	Fiber         *fiber.Server
}

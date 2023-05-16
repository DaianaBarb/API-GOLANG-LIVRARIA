package routes

import (
	rest "api/app/server/rest"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
)

func RegisterRouterLivro(server *fiber.Server, handlers []rest.HandlerLivro) {
	for _, h := range handlers {
		h.RegisterRouterLivro(server)

	}
}

func RegisterRoutes(server *fiber.Server, handlers []rest.Handler) {
	for _, h := range handlers {
		h.RegisterRoute(server)

	}
}

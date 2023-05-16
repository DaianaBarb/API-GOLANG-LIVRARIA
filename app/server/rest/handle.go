package server

import (
	fiberIgnite "github.com/americanas-go/ignite/gofiber/fiber.v2"
)

type HandlerLivro interface {
	RegisterRouterLivro(f *fiberIgnite.Server)
}

type Handler interface {
	RegisterRoute(f *fiberIgnite.Server)
}

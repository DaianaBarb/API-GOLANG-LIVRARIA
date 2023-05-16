package routes

import "api/app/server/rest/routes"

func RegisterRouters(params Params) {

	routes.RegisterRoutes(params.Fiber, params.Handlers)
	routes.RegisterRouterLivro(params.Fiber, params.HandlersLivro)
}

package handler

import (
	"api/app/domain/service"
	util "api/app/domain/util"
	rest "api/app/server/rest"
	"api/app/server/rest/model/request"
	"net/http"

	"github.com/americanas-go/errors"
	f "github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
)

type Livro struct {
	svc       service.LivroServiceInterface
	validator *validator.Validate
}

func NewLivro(service service.LivroServiceInterface, validator *validator.Validate) rest.HandlerLivro {
	return &Livro{
		svc:       service,
		validator: validator,
	}
}

func (h *Livro) RegisterRouterLivro(fServer *f.Server) {
	fServer.App().Add(http.MethodPost, "/partner/:source", h.Save)
	//fServer.App().Add(http.MethodGet, "/partner/:source", h.Find)

}

func (h *Livro) Save(c *fiber.Ctx) error {
	parentCtx := c.Context()
	logger := log.FromContext(parentCtx)

	part := new(request.LivroRequest)
	err := c.BodyParser(&part)
	if err != nil {
		return util.ErrorResponse(c, err)
	}

	body, err := request.NewLivro(c, h.validator)
	if err != nil {
		return util.ErrorResponse(c, err)
	}

	err = h.svc.Save(parentCtx, *body)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
		return util.ErrorResponse(c, err)
	}
	c.Status(http.StatusCreated)
	return nil

}

// func (h *Livro) Find(c *fiber.Ctx) error {
// 	parentCtx := c.Context()
// 	logger := log.FromContext(parentCtx)

// 	param := new(request.LivroRequest)
// 	param.Source = c.Params("source")
// 	param.Source = strings.Trim(param.Source, " ")

// 	source, err := request.NewPartnerParam(c, h.validator)
// 	if err != nil {
// 		return util.ErrorResponse(c, err)
// 	}

// 	partner, err := h.svc.Find(parentCtx, source)
// 	if err != nil {
// 		if errors.IsNotFound(err) {
// 			logger.Debug(errors.ErrorStack(err))
// 			return util.ErrorResponse(c, err)
// 		}

// 		logger.Error(errors.ErrorStack(err))
// 		return util.ErrorResponse(c, err)
// 	}

// 	resp := response.NewPartnerResponseFromDomain(partner)

// 	return c.Status(http.StatusOK).JSON(resp)
// }

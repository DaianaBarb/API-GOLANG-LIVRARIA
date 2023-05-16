package request

import (
	"api/app/domain/util/errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LivroRequest struct {
	Nome      string `json:"name,omitempty"`
	Autor     string `json:"autor,omitempty"`
	Editora   string `json:"editora,omitempty"`
	AnoFabric string `json:"anoFabric,omitempty"`
}

func NewLivro(c *fiber.Ctx, validator *validator.Validate) (*LivroRequest, error) {
	partner := new(LivroRequest)
	if err := c.BodyParser(partner); err != nil {
		return nil, errors.BadRequestf("error to parser body param")
	}

	err := validator.Struct(partner)
	if err != nil {
		return nil, errors.NewBadRequest(err, "error in field validation")
	}

	return partner, nil
}

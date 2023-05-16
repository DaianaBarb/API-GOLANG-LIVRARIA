package util

import (
	"net/http"

	"github.com/americanas-go/errors"
	rest_response "github.com/americanas-go/rest-response"
	fiber "github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, err error) error {
	if errors.IsNotFound(err) {
		return c.Status(http.StatusNotFound).JSON(&rest_response.Error{
			HttpStatusCode: http.StatusNotFound,
			Message:        "resource not found in database",
		})
	}

	if errors.IsBadRequest(err) {
		return c.Status(http.StatusBadRequest).JSON(&rest_response.Error{
			HttpStatusCode: http.StatusBadRequest,
			Message:        err.Error(),
		})
	}

	if errors.IsConflict(err) {
		return c.Status(http.StatusConflict).JSON(&rest_response.Error{
			HttpStatusCode: http.StatusConflict,
			Message:        err.Error(),
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(&rest_response.Error{
		HttpStatusCode: http.StatusInternalServerError,
		Message:        "internal server error",
	})
}

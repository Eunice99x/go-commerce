package api

import (
	"gihub.com/eunice99x/go-commerce/internal/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (a *API) CreateUser(ctx *fiber.Ctx) error {
	var body models.User

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	user, err := a.cfg.UserService.CreateUser(ctx.Context(), &body)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(user)
}

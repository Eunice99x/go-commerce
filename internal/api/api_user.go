package api

import (
	"github.com/eunice99x/go-commerce/internal/database"
	"github.com/eunice99x/go-commerce/internal/models"
	sqlboilermod "github.com/eunice99x/go-commerce/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (a *API) ListUsers(ctx *fiber.Ctx) error {
	// Use the underlying context for SQLBoiler operations
	users, err := sqlboilermod.Users().All(ctx.Context(), database.DB)
	if err != nil {
		// Handle the error appropriately, such as logging or wrapping with Fiber's error
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// Return the list of users as JSON response
	return ctx.JSON(users)
}
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

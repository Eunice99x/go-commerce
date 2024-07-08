package api

import (
	"github.com/eunice99x/go-commerce/internal/database"
	sqlboilermod "github.com/eunice99x/go-commerce/models"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strconv"
)

func (a *API) ListUsers(ctx *fiber.Ctx) error {
	users, err := sqlboilermod.Users().All(ctx.Context(), database.DB)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(users)
}

func (a *API) CreateUser(ctx *fiber.Ctx) error {
	user := sqlboilermod.User{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := user.Insert(ctx.Context(), database.DB, boil.Infer()); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (a *API) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := sqlboilermod.FindUser(ctx.Context(), database.DB, userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if _, err := user.Delete(ctx.Context(), database.DB); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (a *API) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	newUser := sqlboilermod.User{}
	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := sqlboilermod.FindUser(ctx.Context(), database.DB, userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user.Name = newUser.Name
	if _, err := user.Update(ctx.Context(), database.DB, boil.Infer()); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(newUser)
}

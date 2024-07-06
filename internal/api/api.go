package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type API struct {
	app *fiber.App
	cfg *Config
}

type Config struct {
	Port        int
	UserService userService
}

func New(cfg *Config) (*API, error) {
	api := &API{
		cfg: cfg,
		app: fiber.New(fiber.Config{
			BodyLimit: 20971520,
		}),
	}

	if err := api.SetupRoutes(); err != nil {
		return api, nil
	}

	return api, nil
}

func (a *API) Listen() error {
	if err := a.SetupRoutes(); err != nil {
		return fmt.Errorf("setup routes: %v", err)
	}

	if err := a.app.Listen(fmt.Sprintf(":%v", a.cfg.Port)); err != nil {
		return fmt.Errorf("app listen: %v", err)
	}

	return nil
}

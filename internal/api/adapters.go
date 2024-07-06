package api

import (
	"context"
	"gihub.com/eunice99x/go-commerce/internal/models"
)

type userService interface {
	CreateUser(ctx context.Context, filters *models.User) (*models.User, error)
}

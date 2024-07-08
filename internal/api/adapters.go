package api

import (
	"context"
	"github.com/eunice99x/go-commerce/internal/models"
)

type userService interface {
	ListUsers(ctx context.Context, filters *models.User) (*models.User, error)
	CreateUser(ctx context.Context, filters *models.User) (*models.User, error)
}

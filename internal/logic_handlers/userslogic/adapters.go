package userslogic

import (
	"context"
	"github.com/eunice99x/go-commerce/internal/models"
	"github.com/eunice99x/go-commerce/internal/persistence"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . persistor
type persistor interface {
	CreateUser(ctx context.Context, tx persistence.TransactionHandler, user *models.User) (*models.User, error)
}

package userslogic

import (
	"context"
	"github.com/eunice99x/go-commerce/internal/models"
	"github.com/eunice99x/go-commerce/internal/persistence"
	"github.com/sirupsen/logrus"
)

type Config struct {
	TxProvider persistence.TransactionProvider `json:"tx_provider" validate:"required"`
	Logger     *logrus.Entry                   `json:"Logger" validate:"required"`
	Persistor  persistor                       `json:"Persistor" validate:"required"`
}

type Service struct {
	cfg *Config
}

// CreateUser creates a new  user.
func (i *Service) CreateUser(ctx context.Context, params *models.User) (*models.User, error) {

	tx, err := i.cfg.TxProvider.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	user, err := i.cfg.Persistor.CreateUser(ctx, tx, params)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

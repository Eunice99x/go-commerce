package persistence

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// TransactionProvider
//
//counterfeiter:generate . TransactionProvider
type TransactionProvider interface {
	Db(ctx context.Context) (TransactionHandler, error)
	Tx(ctx context.Context) (TransactionHandler, error)
}

// TransactionHandler contains our generic exposed methods
// to present to the service/biz layer.
//
//counterfeiter:generate . TransactionHandler
type TransactionHandler interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context)
	GetCtxExecutor(i interface{}) (boil.ContextExecutor, error)
}
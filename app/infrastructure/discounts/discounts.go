package discounts

import (
	"context"
	"errors"
	"math/rand"

	"github.com/AnhTaFP/go-error-handling/app/domain/discounts"
	domainerrors "github.com/AnhTaFP/go-error-handling/app/domain/errors"
)

type DB struct {
	host     string
	username string
	password string
}

func NewDB(host string, username string, password string) *DB {
	return &DB{
		host:     host,
		username: username,
		password: password,
	}
}

func (db *DB) List(ctx context.Context, customer string) ([]discounts.Discount, error) {
	ds, err := db.query(ctx)
	if err != nil {
		// the usual way to wrap error
		// return nil, fmt.Errorf("error querying db: %w", err)
		// instead, we should do
		return nil, domainerrors.Wrap(err, map[string]interface{}{
			"category": "infrastructure",
			"service":  "discounts db",
		}, "cannot get discounts list for customer %s", customer)
	}

	return ds, nil
}

func (db *DB) query(ctx context.Context) ([]discounts.Discount, error) {
	i := rand.Intn(3)

	// let's pretend that this actually calls an external database like Amazon RDS or DynamoDB
	// assume that arbitrary errors like too many connections, or lost connections can happen
	switch i {
	case 0:
		return nil, errTooManyConnections
	case 1:
		return nil, errLostConnection
	default:
		return []discounts.Discount{
			{
				ID:    "1",
				Title: "discount #1",
				Value: 5.5,
			},
			{
				ID:    "2",
				Title: "discount #2",
				Value: 6.5,
			},
		}, nil
	}
}

var (
	errTooManyConnections = errors.New("too many connections")
	errLostConnection     = errors.New("lost connection")
)

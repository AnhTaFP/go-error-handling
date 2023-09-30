package flag

import (
	"context"
	"errors"
	"math/rand"

	domainerrors "github.com/AnhTaFP/go-error-handling/app/domain/errors"
	"github.com/AnhTaFP/go-error-handling/app/domain/optimization"
)

type Service struct {
	url string
}

func NewService(url string) *Service {
	return &Service{
		url: url,
	}
}

func (s *Service) GetFlag(ctx context.Context, flag string) (*optimization.Flag, error) {
	f, err := s.get(ctx, flag)
	if err != nil {
		return nil, domainerrors.Wrap(err, map[string]interface{}{
			"category": "infrastructure",
			"service":  "flag",
		}, "cannot get flag value for %", flag)
	}

	return f, nil
}

func (s *Service) get(ctx context.Context, flag string) (*optimization.Flag, error) {
	i := rand.Intn(4)

	// let's pretend that this actually calls an external API to verify token
	// arbitrary errors like i/o timeout and context cancelled can happen as well
	switch i {
	case 0:
		return &optimization.Flag{Enabled: false}, nil
	case 1:
		return nil, errIOTimeout
	case 2:
		return nil, errContextCancelled
	default:
		return &optimization.Flag{Enabled: false}, nil
	}
}

var (
	errIOTimeout        = errors.New("i/o timeout")
	errContextCancelled = errors.New("context cancelled")
)

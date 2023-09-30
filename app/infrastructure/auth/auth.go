package auth

import (
	"context"
	"errors"
	"math/rand"

	domainerrors "github.com/AnhTaFP/go-error-handling/app/domain/errors"
)

type Service struct {
	url string
}

func NewService(url string) *Service {
	return &Service{url: url}
}

func (s *Service) VerifyToken(ctx context.Context, token string) error {
	tk, err := s.verify(ctx, token)
	if err != nil {
		// the usual way to wrap error
		// return fmt.Errorf("error verifying token: %w", err)
		// instead we should do
		return domainerrors.Wrap(err, map[string]interface{}{
			"category": "infrastructure",
			"service":  "auth",
		}, "cannot verify customer token")
	}

	if !tk.valid {
		return domainerrors.Wrap(nil, map[string]interface{}{
			"category": "infrastructure",
			"service":  "auth",
		}, "invalid customer token")
	}

	return nil
}

func (s *Service) verify(ctx context.Context, token string) (*customerToken, error) {
	i := rand.Intn(4)

	// let's pretend that this actually calls an external API to verify token
	// assume that when token service rejects the token, we get a customerToken with valid = false
	// assume that when token service accepts the token, we get a customerToken with valid = true
	// arbitrary errors like i/o timeout and context cancelled can happen as well
	switch i {
	case 0:
		return &customerToken{valid: false}, nil
	case 1:
		return nil, errIOTimeout
	case 2:
		return nil, errContextCancelled
	default:
		return &customerToken{valid: true}, nil
	}
}

var (
	errIOTimeout        = errors.New("i/o timeout")
	errContextCancelled = errors.New("context cancelled")
)

type customerToken struct {
	valid bool
	// other fields
}

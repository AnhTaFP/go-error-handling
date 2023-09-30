package usecase

import (
	"context"

	"github.com/AnhTaFP/go-error-handling/app/domain/discounts"
)

// ListDiscounts is a hypothetical use case: user will have their token verified first,
// if successful, the list of discounts for the user will be queried from the database,
// then the list of discount will be optimized.
type ListDiscounts struct {
	authService        AuthService
	discountRepository DiscountRepository
	optimizer          Optimizer
}

func NewListDiscounts(
	authService AuthService,
	discountRepository DiscountRepository,
	optimizer Optimizer,
) *ListDiscounts {
	return &ListDiscounts{
		authService:        authService,
		discountRepository: discountRepository,
		optimizer:          optimizer,
	}
}

func (ld *ListDiscounts) List(ctx context.Context, token string, customer string) ([]discounts.Discount, error) {
	if err := ld.authService.VerifyToken(ctx, token); err != nil {
		return nil, err
	}

	ds, err := ld.discountRepository.List(ctx, customer)
	if err != nil {
		return nil, err
	}

	ds = ld.optimizer.Optimize(ctx, ds)

	return ds, nil
}

type AuthService interface {
	VerifyToken(ctx context.Context, token string) error
}

type DiscountRepository interface {
	List(ctx context.Context, customer string) ([]discounts.Discount, error)
}

type Optimizer interface {
	Optimize(ctx context.Context, ds []discounts.Discount) []discounts.Discount
}

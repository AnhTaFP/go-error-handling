package optimization

import (
	"context"

	"github.com/AnhTaFP/go-error-handling/app/domain/discounts"
	domainerrors "github.com/AnhTaFP/go-error-handling/app/domain/errors"

	"github.com/sirupsen/logrus"
)

type DiscountOptimizer struct {
	entry       *logrus.Entry
	flagService FlagService
}

func NewDiscountOptimizer(flagService FlagService) *DiscountOptimizer {
	return &DiscountOptimizer{flagService: flagService}
}

func (o *DiscountOptimizer) Optimize(ctx context.Context, ds []discounts.Discount) []discounts.Discount {
	flag, err := o.flagService.GetFlag(ctx, "optimizer_enabled")
	if err != nil {
		o.logError(err)
		return ds
	}

	if flag.Enabled {
		ds = o.optimize(ctx, ds)
	}

	return ds
}

func (o *DiscountOptimizer) logError(err error) {
	domainErr, ok := err.(domainerrors.Error)
	if ok {
		o.entry.WithError(domainErr).WithFields(logrus.Fields{
			"category": domainErr.Misc["category"],
			"service":  domainErr.Misc["service"],
			"scope":    "optimization",
		}).Error(domainErr.Error())
	} else {
		o.entry.WithError(err).WithFields(logrus.Fields{
			"scope": "optimization",
		}).Error("encounter generic error")
	}
}

func (o *DiscountOptimizer) optimize(ctx context.Context, ds []discounts.Discount) []discounts.Discount {
	// optimization logic goes here
	return ds
}

type FlagService interface {
	GetFlag(ctx context.Context, flag string) (*Flag, error)
}

type Flag struct {
	Name    string
	Enabled bool
}

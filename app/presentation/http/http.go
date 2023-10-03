package http

import (
	"encoding/json"
	"net/http"

	domainerrors "github.com/AnhTaFP/go-error-handling/app/domain/errors"
	"github.com/AnhTaFP/go-error-handling/app/domain/optimization"
	"github.com/AnhTaFP/go-error-handling/app/infrastructure/auth"
	"github.com/AnhTaFP/go-error-handling/app/infrastructure/discounts"
	"github.com/AnhTaFP/go-error-handling/app/usecase"
	"github.com/sirupsen/logrus"
)

func ListDiscounts(
	entry *logrus.Entry,
	auth *auth.Service,
	discountsRepository *discounts.DB,
	optimizer *optimization.DiscountOptimizer,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		listDiscounts := usecase.NewListDiscounts(auth, discountsRepository, optimizer)
		ds, err := listDiscounts.List(r.Context(), "token-xyz", "customer-abcdef")
		if err != nil {
			logError(entry, err)

			// if err is domainerrors.Error, then return only the friendly message to the clients.
			if dErr, ok := err.(*domainerrors.Error); ok {
				respondError(w, dErr.FriendlyMessage)
				return
			}

			// otherwise we return a generic message to the clients.
			respondError(w, "internal server error")
			return
		}

		var res response
		for _, d := range ds {
			res.Discounts = append(res.Discounts, discount{
				ID:    d.ID,
				Title: d.Title,
				Value: d.Value,
			})
		}

		b, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func respondError(w http.ResponseWriter, msg string) {
	var e struct {
		Error string `json:"error"`
	}

	e.Error = msg
	b, _ := json.Marshal(e)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}

func logError(entry *logrus.Entry, err error) {
	domainErr, ok := err.(*domainerrors.Error)
	if ok {
		entry.WithError(domainErr).WithFields(logrus.Fields{
			"category": domainErr.Misc["category"],
			"service":  domainErr.Misc["service"],
		}).Error(domainErr.Error())
	} else {
		entry.WithError(err).Error("encounter generic error")
	}
}

type response struct {
	Discounts []discount `json:"discounts"`
}

type discount struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Value float32 `json:"value"`
}

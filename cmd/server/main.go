package main

import (
	"log"
	sdthttp "net/http"

	"github.com/AnhTaFP/go-error-handling/app/domain/optimization"
	"github.com/AnhTaFP/go-error-handling/app/infrastructure/auth"
	"github.com/AnhTaFP/go-error-handling/app/infrastructure/discounts"
	"github.com/AnhTaFP/go-error-handling/app/infrastructure/flag"
	"github.com/AnhTaFP/go-error-handling/app/presentation/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	authService := auth.NewService("auth service URL")
	discountsRepository := discounts.NewDB("db-host", "db-username", "db-password")
	flagService := flag.NewService("flag service URL")
	optimizer := optimization.NewDiscountOptimizer(flagService)

	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	entry := logrus.NewEntry(logger)

	r := mux.NewRouter()
	r.HandleFunc("/discounts", http.ListDiscounts(entry, authService, discountsRepository, optimizer))

	if err := sdthttp.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatalln("error with http server:", err.Error())
	}
}

package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/LeoUraltsev/HauseService/internal/config"
	"github.com/LeoUraltsev/HauseService/internal/gen"
	"github.com/LeoUraltsev/HauseService/internal/handlers"
	"github.com/LeoUraltsev/HauseService/internal/service"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/LeoUraltsev/HauseService/internal/storage/postgres"
)

func Run(log *slog.Logger, cfg *config.Config) error {

	db, err := postgres.New(context.Background(), cfg.PostgresURLConnection, log)
	if err != nil {
		return err
	}
	defer db.Close()

	authService := service.NewAuthService(db)

	r := gen.HandlerWithOptions(
		handlers.New(db, db, authService),
		gen.ChiServerOptions{
			Middlewares: []gen.MiddlewareFunc{
				middleware.RequestID,
			},
		},
	)

	http.ListenAndServe("localhost:10000", r)
	return nil
}

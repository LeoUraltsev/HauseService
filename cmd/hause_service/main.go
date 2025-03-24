package main

import (
	"log/slog"
	"os"

	"github.com/LeoUraltsev/HauseService/internal/config"
)

const (
	local = "local"
	prod  = "prod"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error(err.Error())
	}

	log := initLogger(cfg.Env)

	log.Info("startup app", slog.String("env", cfg.Env))

}

func initLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case prod:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return log
}

package apiserver

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

func Start(config *Config, ctx context.Context) error {
	db, err := newDB(config.DatabaseURL, ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	logger := logrus.New()

	srv := newServer(logger, config.BindAddr)

	go func(*logrus.Logger) {
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Fatal(err)
			}
		}
	}(logger)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error(err)
	}

	logger.Error("performing graceful shutdown")
	return nil
}

func newDB(databaseURL string, ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

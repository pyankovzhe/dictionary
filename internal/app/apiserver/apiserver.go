package apiserver

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pyankovzhe/dictionary/internal/app/store/sqlstore"
	"github.com/sirupsen/logrus"
)

func Start(config *Config, ctx context.Context) error {
	db, err := sqlstore.NewDB("pgx", config.DatabaseURL, ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	logger := logrus.New()
	store := sqlstore.New(db)

	srv := newServer(logger, store, config.BindAddr)

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

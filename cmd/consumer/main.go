package main

import (
	"context"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/pyankovzhe/dictionary/internal/app/store/sqlstore"
	"github.com/pyankovzhe/dictionary/internal/consumer"
	"github.com/pyankovzhe/dictionary/internal/consumer/kafkaconsumer"
	"github.com/sirupsen/logrus"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/consumer.toml", "path to consumer config file")
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()
	config := consumer.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	db, err := sqlstore.NewDB("pgx", config.DatabaseURL, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger := logrus.New()
	store := sqlstore.New(db)

	consumer, err := kafkaconsumer.New(ctx, logger, store, config)
	if err != nil {
		logger.Fatal(err)
	}

	consumer.Consume()
}

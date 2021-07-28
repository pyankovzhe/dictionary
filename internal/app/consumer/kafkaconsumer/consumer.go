package kafkaconsumer

import (
	"context"
	"encoding/json"

	"github.com/pyankovzhe/dictionary/internal/app/model"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	reader *kafka.Reader
	ctx    context.Context
	logger *logrus.Logger
}

func New(ctx context.Context, logger *logrus.Logger, address string, topic string, partition int) (*Consumer, error) {
	config := kafka.ReaderConfig{
		Brokers:   []string{address},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		GroupID:   "dictionary-group",
	}

	r := kafka.NewReader(config)
	consumer := &Consumer{reader: r, ctx: ctx, logger: logger}
	consumer.logger.Info("Kafka consumer initialized")

	return consumer, nil
}

func (c *Consumer) Close() {
	c.reader.Close()
}

func (c *Consumer) Consume() {
	c.logger.Info("Kafka starts to consume messages...")

	for {
		m, err := c.reader.ReadMessage(c.ctx)
		if err != nil {
			c.logger.Error("Fail to read kafka message", err)
			break
		}

		acc := &model.Account{}
		if err := json.Unmarshal(m.Value, acc); err != nil {
			c.logger.Errorf("Fail to unmarshal kafka message with offset %d: %s. Err: %s", m.Offset, m.Value, err)
		} else {
			c.logger.Logf(4, "message at offset %d: %v", m.Offset, acc)
		}
	}
}

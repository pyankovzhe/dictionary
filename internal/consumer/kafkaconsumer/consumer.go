package kafkaconsumer

import (
	"context"

	"github.com/pyankovzhe/dictionary/internal/app/store"
	"github.com/pyankovzhe/dictionary/internal/consumer"
	pb "github.com/pyankovzhe/dictionary/pkg/proto/v1/eventpb"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type Consumer struct {
	reader *kafka.Reader
	ctx    context.Context
	logger *logrus.Logger
	store  store.Store
}

func New(ctx context.Context, logger *logrus.Logger, store store.Store, config *consumer.Config) (*Consumer, error) {
	readerConfig := kafka.ReaderConfig{
		Brokers:   []string{config.KafkaURL},
		Topic:     config.Topic,
		Partition: config.Partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		GroupID:   config.GroupID,
	}

	r := kafka.NewReader(readerConfig)
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

		event := &pb.AccountEvent{}

		if err := proto.Unmarshal(m.Value, event); err != nil {
			c.logger.Errorf("Fail to unmarshal kafka message with offset %d: %s. Err: %s", m.Offset, m.Value, err)
		} else {
			c.logger.Logf(4, "message at offset %d: %v", m.Offset, event)
		}

		switch event.Kind {
		case pb.EventKind_CREATED:
			c.logger.Info("create account")
		default:
			c.logger.Logf(4, "nooop %v", event.Kind)
		}
	}
}

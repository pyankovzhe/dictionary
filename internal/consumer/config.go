package consumer

type Config struct {
	KafkaURL    string `toml:"kafka_url"`
	Topic       string `toml:"topic"`
	Partition   int    `toml:"partition"`
	DatabaseURL string `toml:"database_url"`
	GroupID     string `toml:"group_id"`
}

func NewConfig() *Config {
	return &Config{
		KafkaURL:  "localhost:9093",
		Topic:     "accounts",
		Partition: 0,
		GroupID:   "dictionary-group",
	}
}

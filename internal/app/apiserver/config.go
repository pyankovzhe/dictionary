package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
	KafkaURL    string `toml:"kafka_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
	}
}

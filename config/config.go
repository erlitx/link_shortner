package config

import (
	"fmt"

	"github.com/erlitx/link_shortner/internal/adapter/kafka_producer"
	minioadapter "github.com/erlitx/link_shortner/internal/adapter/miniio"
	"github.com/erlitx/link_shortner/internal/controller/kafka_consumer"
	"github.com/erlitx/link_shortner/internal/controller/worker"
	"github.com/erlitx/link_shortner/pkg/httpserver"
	"github.com/erlitx/link_shortner/pkg/postgres"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App           App
	HTTP          httpserver.Config
	Postgres      postgres.Config
	KafkaProducer kafka_producer.Config
	KafkaConsumer kafka_consumer.Config
	ProduceWorker worker.ProduceConfig
	MiniIo        minioadapter.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}

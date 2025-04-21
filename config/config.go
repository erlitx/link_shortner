package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/httpserver"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/postgres"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/wb"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/adapter/kafka"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App      App
	HTTP     httpserver.Config
	Postgres postgres.Config
	WB       wb.Config
	KafkaProducer kafka.Config 
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

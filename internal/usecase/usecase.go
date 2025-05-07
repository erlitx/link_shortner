package usecase

import (
	"context"

	"github.com/segmentio/kafka-go"

	"github.com/erlitx/link_shortner/internal/domain"
	"github.com/erlitx/link_shortner/internal/dto"
)

type Cache interface {
	Set(url domain.URL) error
	Get(input dto.GetURLInput) (domain.URL, bool)
}

type Postgres interface {
	CreateShortURL(ctx context.Context, p domain.URL) error
	ResolveShortURL(ctx context.Context, short string) (string, error)
}

type KafkaProducer interface {
	Produce(ctx context.Context, msgs ...kafka.Message) error
}

type Storage interface {
	SaveFile(ctx context.Context, key string, data []byte) error 
}

type UseCase struct {
	cache          Cache
	postgres       Postgres
	kafka_producer KafkaProducer
	storage        Storage
}

func New(cache Cache, p Postgres, kafProd KafkaProducer, s Storage) *UseCase {
	return &UseCase{
		cache:    cache,
		postgres: p,
		kafka_producer: kafProd,
		storage: s,
	}
}

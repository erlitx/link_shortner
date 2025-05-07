package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

func (u *UseCase) SendShortURLtoKafka(url string) error {
	ctx := context.Background()
	log.Info().Msg("---------LOGGED-------")

	msg := kafka.Message{
		Topic: "qr_code_task",
		Value: []byte(url),
	}

	err := u.kafka_producer.Produce(ctx, msg)

	if err != nil {
		log.Error().Err(err).Msg("produce worker: some work failed")
	}

	return nil
}

package usecase

import (
	"context"
	"fmt"

	"github.com/erlitx/link_shortner/internal/domain"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func (u *UseCase) Produce(ctx context.Context, msgCount int) error {
	msgs := make([]domain.Message, 0, msgCount)

	for range msgCount {
		msgs = append(msgs, domain.Message{
			Topic: "qr_code_task",
			Key:   []byte(uuid.New().String()),
			Value: []byte(uuid.New().String()),
		})
	}

	// Send to Kafka
	err := u.kafka_producer.Produce(ctx, toKafkaMessages(msgs)...)
	if err != nil {
		return fmt.Errorf("u.kafka.Produce: %w", err)
	}

	return nil
}

func toKafkaMessages(msgs []domain.Message) []kafka.Message {
	kafkaMsgs := make([]kafka.Message, 0, len(msgs))

	for _, msg := range msgs {
		kafkaMsgs = append(kafkaMsgs, kafka.Message{
			Topic: msg.Topic,
			Key:   msg.Key,
			Value: msg.Value,
		})
	}
	return kafkaMsgs
}

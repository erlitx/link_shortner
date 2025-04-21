package kafka

import (
	"context"
	"fmt"
	"hash/fnv"

	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/metrics"

	"github.com/segmentio/kafka-go"
)

type Config struct {
	Addr  []string `envconfig:"KAFKA_WRITER_ADDR" required:"true"`
	Topic string   `envconfig:"KAFKA_WRITER_TOPIC" default:"mnepryakhin-my-app-topic"`
}

type Producer struct {
	writer *kafka.Writer
	metrics *metrics.Entity
}

func NewProducer(c Config, metrics *metrics.Entity) *Producer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(c.Addr...),
		Topic:        c.Topic,
		Balancer:     &kafka.Hash{Hasher: fnv.New32a()},
		RequiredAcks: kafka.RequireAll,
		//ErrorLogger:  logger.ErrorLogger(),
		// Async:        true,
	}

	return &Producer{
		writer: w,
		metrics: metrics,
	}
}

func (p *Producer) Produce(ctx context.Context, msgs ...kafka.Message) error {
	const produce = "produce"

	//defer p.metrics.Duration(produce, time.Now())

	err := p.writer.WriteMessages(ctx, msgs...)
	if err != nil {
		p.metrics.TotalAdd(produce, metrics.Error, len(msgs))

		return fmt.Errorf("p.writer.WriteMessages: %w", err)
	}

	p.metrics.TotalAdd(produce, metrics.Ok, len(msgs))

	return nil
}

func (p *Producer) Close() {
	err := p.writer.Close()
	if err != nil {
		log.Error().Err(err).Msg("kafka producer: p.writer.Close")
	}

	log.Info().Msg("kafka producer: closed")
}

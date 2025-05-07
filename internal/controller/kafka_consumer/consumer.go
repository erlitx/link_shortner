package kafka_consumer

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/erlitx/link_shortner/internal/usecase"
	"github.com/erlitx/link_shortner/pkg/logger"
	"github.com/erlitx/link_shortner/pkg/metrics"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type Config struct {
	Addr     []string `envconfig:"KAFKA_CONSUMER_ADDR" required:"true"`
	Topic    string   `envconfig:"KAFKA_CONSUMER_TOPIC" default:"qr_code_task"`
	Group    string   `envconfig:"KAFKA_CONSUMER_GROUP" default:"qr_code_group"`
	Disabled bool     `envconfig:"KAFKA_CONSUMER_DISABLED"`
}

type Consumer struct {
	config  Config
	reader  *kafka.Reader
	usecase *usecase.UseCase
	metrics *metrics.Entity
	stop    context.CancelFunc
	done    chan struct{}
}

func New(cfg Config, metrics *metrics.Entity, uc *usecase.UseCase) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          cfg.Addr,
		Topic:            cfg.Topic,
		GroupID:          cfg.Group,
		ErrorLogger:      logger.ErrorLogger(),
		ReadBatchTimeout: time.Second,
		// CommitInterval: time.Second,
	})

	ctx, stop := context.WithCancel(context.Background())

	c := &Consumer{
		config:  cfg,
		reader:  r,
		usecase: uc,
		metrics: metrics,
		stop:    stop,
		done:    make(chan struct{}),
	}

	if c.config.Disabled {
		log.Info().Msg("kafka consumer: disabled")

		return c
	}

	go c.run(ctx)

	return c
}

func (c *Consumer) run(ctx context.Context) {
	const consume = "get_url_kafka"

	log.Info().Msg("kafka consumer: started")

	for {
		now := time.Now()

		m, err := c.reader.FetchMessage(ctx) // Read Message
		if err != nil {
			log.Error().Err(err).Msg("kafka consumer: FetchMessage")

			if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
				break
			}
		}
		// Generate QR code and save to S3
		go func() {
			err = c.usecase.GenerateQRCode(ctx, m) 
			if err != nil {
				c.metrics.Total(consume, metrics.Error)
				log.Error().Err(err).Msg("kafka consumer: some work failed")
			}
		}()

		if err = c.reader.CommitMessages(ctx, m); err != nil {
			c.metrics.Total(consume, metrics.Error)
			log.Error().Err(err).Msg("kafka consumer: CommitMessages")
		}

		c.metrics.Duration(consume, now)
		c.metrics.Total(consume, metrics.Ok)
	}

	close(c.done)
}

func (c *Consumer) Close() {
	if c.config.Disabled {
		return
	}

	log.Info().Msg("kafka consumer: closing")

	c.stop()

	if err := c.reader.Close(); err != nil {
		log.Error().Err(err).Msg("kafka consumer: reader.Close")
	}

	<-c.done

	log.Info().Msg("kafka consumer: closed")
}

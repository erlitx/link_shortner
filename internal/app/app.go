package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/erlitx/link_shortner/config"
	"github.com/erlitx/link_shortner/internal/adapter/cache"
	"github.com/erlitx/link_shortner/internal/adapter/kafka_producer"
	"github.com/erlitx/link_shortner/internal/controller/kafka_consumer"
	postgresAdapter "github.com/erlitx/link_shortner/internal/adapter/postgres"
	"github.com/erlitx/link_shortner/internal/adapter/miniio"

	"github.com/erlitx/link_shortner/internal/controller/http"
	//"github.com/erlitx/link_shortner/internal/controller/worker"
	"github.com/erlitx/link_shortner/internal/usecase"
	"github.com/erlitx/link_shortner/pkg/httpserver"
	postgresPkg "github.com/erlitx/link_shortner/pkg/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/erlitx/link_shortner/pkg/metrics"
)

type Dependencies struct {
	// Adapters
	Postgres *postgresPkg.Pool

}

func Run(ctx context.Context, c config.Config) (err error) {
	var deps Dependencies

	entityMetrics := metrics.NewEntity("url_shortner")

	// CREATING ADAPTERS
	// 1. Postgres
	deps.Postgres, err = postgresPkg.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}
	defer deps.Postgres.Close()

	// 2. Cache
	cacheAdapter := cache.New(2)
	pgPool := postgresAdapter.New(deps.Postgres.Pool)

	// 3. Kafka Producer
	kafkaProducer := kafka_producer.NewProducer(c.KafkaProducer, entityMetrics)

	// MINIIO
	minioClient, err := minioadapter.New(c.MiniIo)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init MinIO")
	}


	// USECASE
	//Passing adapters
	uc := usecase.New(cacheAdapter, pgPool, kafkaProducer, minioClient)

	// 4. Kafka consumer
	kafkaConsumer := kafka_consumer.New(c.KafkaConsumer, entityMetrics, uc)

	// HTTP
	router := chi.NewRouter()
	http.ProfileRouter(router, uc)
	httpServer := httpserver.New(router, "3000")

	// Produce worker
	// produceWorker := worker.NewProduceWorker(c.ProduceWorker, uc)

	// log.Info().Msg("App started!")

	// STOPPING
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	log.Info().Msg("App got signal to stop")

	// Contollers
	kafkaConsumer.Close()
	httpServer.Close()
	//produceWorker.Stop()

	// Adapters
	kafkaProducer.Close()
	deps.Postgres.Close()

	log.Info().Msg("App stopped!")

	return nil
}


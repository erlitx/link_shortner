package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/config"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/adapter/cache"
	postgresAdapter "gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/adapter/postgres"
	wbadapter "gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/adapter/wb"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/controller/http"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/usecase"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/httpserver"
	postgresPkg "gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/postgres"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/wb"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/adapter/kafka"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/metrics"
)

type Dependencies struct {
	// Adapters
	Postgres *postgresPkg.Pool
	WB *wb.Client
	// KafkaWriter *kafka_writer.Writer
	// Redis       *redis.Client

	// Controllers
	// RouterHTTP  *chi.Mux
	// KafkaReader *kafka_reader.Reader
}

func Run(ctx context.Context, c config.Config) (err error) {
	var deps Dependencies

	// Adapters
	// Postgres
	deps.Postgres, err = postgresPkg.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}
	defer deps.Postgres.Close()

	entityMetrics := metrics.NewEntity("my-app")
	kafkaProducer := kafka.NewProducer(c.KafkaProducer, entityMetrics)

	// Cache
	cacheAdapter := cache.New()
	pgPool := postgresAdapter.New(deps.Postgres.Pool)

	// WB Client
	deps.WB = wb.New(c.WB)
	wbClient := wbadapter.New(deps.WB)

	uc := usecase.New(cacheAdapter, pgPool, wbClient)

	router := chi.NewRouter()

	http.ProfileRouter(router, uc)

	httpServer := httpserver.New(router, "3000")
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *httpserver.Server) {
	log.Info().Msg("App started!")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		log.Info().Msg("App got signal: " + i.String())
	case err := <-httpServer.Notify():
		log.Error().Err(err).Msg("App got notify: httpServer.Notify")
	}

	log.Info().Msg("App is stopping...")
}

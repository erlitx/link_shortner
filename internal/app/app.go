package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/erlitx/link_shortner/config"
	"github.com/erlitx/link_shortner/internal/adapter/cache"
	postgresAdapter "github.com/erlitx/link_shortner/internal/adapter/postgres"
	"github.com/erlitx/link_shortner/internal/controller/http"
	"github.com/erlitx/link_shortner/internal/usecase"
	"github.com/erlitx/link_shortner/pkg/httpserver"
	postgresPkg "github.com/erlitx/link_shortner/pkg/postgres"
)

type Dependencies struct {
	// Adapters
	Postgres *postgresPkg.Pool

}

func Run(ctx context.Context, c config.Config) (err error) {
	var deps Dependencies

	// CREATING ADAPTERS
		// Postgres
	deps.Postgres, err = postgresPkg.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}
	defer deps.Postgres.Close()

		// Cache
	cacheAdapter := cache.New(2)
	pgPool := postgresAdapter.New(deps.Postgres.Pool)


	// CREATING USECASE
		//Passing adapters
	uc := usecase.New(cacheAdapter, pgPool)

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

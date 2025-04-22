package main

import (
	"context"

	"github.com/erlitx/link_shortner/config"
	"github.com/erlitx/link_shortner/internal/app"
	"github.com/erlitx/link_shortner/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init(logger.Config{
		AppName:       "LINK-SHORTNER",
		AppVersion:    "v0.1.0",
		Level:         "debug",
		PrettyConsole: true,
	})

	ctx := context.Background()

	config, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	err = app.Run(ctx, config)
	if err != nil {
		panic(err)
	}

}

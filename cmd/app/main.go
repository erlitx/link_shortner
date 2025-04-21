package main

import (
	"context"

	"github.com/pkg/profile"
	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/config"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/app"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/pkg/logger"
)

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	logger.Init(logger.Config{
		AppName:       "my-app",
		AppVersion:    "v0.1.0",
		Level:         "debug",
		PrettyConsole: true,
	})

	ctx := context.Background()

	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	err = app.Run(ctx, c)
	if err != nil {
		panic(err)
	}

}

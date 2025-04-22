package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/rs/zerolog/log"
	"github.com/erlitx/link_shortner/internal/domain"
)


func (postgres *Postgres) CreateShortURL(ctx context.Context, url domain.URL) (err error) {
	fmt.Println("Postgres - create profile")

	dialect := goqu.Dialect("postgres")
	record := goqu.Record{"rawurl": url.RawURL, "shorturl": url.ShortURL}

	sql, _, _ := dialect.Insert("public.urls").Rows(record).ToSQL()

	log.Info().Msg(sql)

	_, err = postgres.pool.Exec(ctx, sql)
	
	if err != nil {
		log.Error().Err(err).Msg("Error")
	}

	return nil
}
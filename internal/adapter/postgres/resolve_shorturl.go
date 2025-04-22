package postgres

import (
	"context"
	"fmt"
)

func (postgres *Postgres) ResolveShortURL(ctx context.Context, short string) (shortUrl string, err error) {
	var rawURL string

	sqlStr := "SELECT rawurl FROM public.urls WHERE shorturl = $1 LIMIT 1"

	err = postgres.pool.QueryRow(ctx, sqlStr, short).Scan(&rawURL)
	if err != nil {
		return "", fmt.Errorf("ResolveShortURL: %w", err)
	}

	return rawURL, nil
}

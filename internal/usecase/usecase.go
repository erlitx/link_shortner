package usecase

import (
	"context"

	"github.com/erlitx/link_shortner/internal/domain"
	"github.com/erlitx/link_shortner/internal/dto"
)

type Cache interface {
	Set(url domain.URL)
	Get(input dto.GetURLInput) (domain.URL, bool)
}

type Postgres interface {
	CreateShortURL(ctx context.Context, p domain.URL) error
	ResolveShortURL(ctx context.Context, short string) (string, error)
}

type UseCase struct {
	cache    Cache
	postgres Postgres
}

func New(cache Cache, p Postgres) *UseCase {
	return &UseCase{
		cache:    cache,
		postgres: p,
	}
}

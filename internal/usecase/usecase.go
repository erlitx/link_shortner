package usecase

import (
	"context"

	"github.com/erlitx/link_shortner/internal/domain"
	"github.com/google/uuid"
)

type Cache interface {
	Add(key uuid.UUID, profile domain.Profile)
	Get(key uuid.UUID) (domain.Profile, error)
	Update(key uuid.UUID, profile domain.Profile)
	Delete(key uuid.UUID)
}

type Postgres interface {
	CreateShortURL(ctx context.Context, p domain.URL) error
	ResolveShortURL(ctx context.Context, short string) (string, error)
	GetProfile(domain.Profile) error
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

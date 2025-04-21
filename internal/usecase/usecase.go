package usecase

import (
	"context"

	"github.com/google/uuid"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/domain"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/dto"
)

type Cache interface {
	Add(key uuid.UUID, profile domain.Profile)
	Get(key uuid.UUID) (domain.Profile, error)
	Update(key uuid.UUID, profile domain.Profile)
	Delete(key uuid.UUID)
}

type Postgres interface {
	CreateProfile(ctx context.Context, p domain.Profile) error
	GetProfile(domain.Profile) error
	UploadWBOrders(ctx context.Context, p []domain.WBorder) (err error) 
}

type WB interface {
	GetWBOrders(ctx context.Context, input dto.GetWBOrdersInput) ([]dto.GetWBordersOutput, error)
}

type UseCase struct {
	cache    Cache
	postgres Postgres
	wb       WB
}

func New(cache Cache, p Postgres, wb WB) *UseCase {
	return &UseCase{
		cache:    cache,
		postgres: p,
		wb:       wb,
	}
}

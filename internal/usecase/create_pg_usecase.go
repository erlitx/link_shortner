package usecase

import (
	"context"

	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/domain"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/dto"
)

func (u *UseCase) CreatePGProfile(ctx context.Context) (dto.CreateProfileOutput, error) {
	u.postgres.CreateProfile(ctx, domain.Profile{})
	return dto.CreateProfileOutput{}, nil
}

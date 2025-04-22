package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/erlitx/link_shortner/internal/domain"
	"github.com/erlitx/link_shortner/internal/dto"
)

func (u *UseCase) UpdateProfile(input dto.UpdateProfileInput) error {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}

	profile, err := domain.NewProfile(input.Name, input.Age, id)
	if err != nil {
		return fmt.Errorf("domain.NewProfile: %w", err)
	}

	u.cache.Update(id, profile)

	return nil
}

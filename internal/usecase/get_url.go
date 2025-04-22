package usecase

import (
	"context"
	"fmt"

	"github.com/erlitx/link_shortner/internal/dto"
)

func (u *UseCase) ResolveShortURL(ctx context.Context, input dto.GetURLInput) (dto.GetURLOutput, error) {
	var output dto.GetURLOutput
	fmt.Println("----", input)
	rawUrl, err := u.postgres.ResolveShortURL(ctx, input.ShortUrl)
	if err != nil {
		return output, fmt.Errorf("shortUrl.Resolve: %w", err)
	}


	// id, err := uuid.Parse(input.ID)
	// if err != nil {
	// 	return output, fmt.Errorf("uuid.Parse: %w", err)
	// }

	// profile, err := u.cache.Get(id)
	// if err != nil {
	// 	return output, fmt.Errorf("u.cache.Get: %w", err)
	// }

	output.RedirectURL = rawUrl
	return output, nil
}

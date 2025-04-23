package usecase

import (
	"context"
	"fmt"

	"github.com/erlitx/link_shortner/internal/domain"
	"github.com/erlitx/link_shortner/internal/dto"
)

func (u *UseCase) ResolveShortURL(ctx context.Context, input dto.GetURLInput) (dto.GetURLOutput, error) {
	var output dto.GetURLOutput

	// Check cache
	url, ok := u.cache.Get(input)
	if ok {
		return dto.GetURLOutput{RedirectURL: string(url.RawURL)}, nil
	}

	fmt.Println("POSTGRES CHECKED")

	rawUrl, err := u.postgres.ResolveShortURL(ctx, input.ShortUrl)

	// Save to cache
	domainUrl := domain.URL{RawURL: domain.RawURL(rawUrl), ShortURL: domain.ShortURL(input.ShortUrl)}
	u.cache.Set(domainUrl)

	if err != nil {
		return output, fmt.Errorf("shortUrl.Resolve: %w", err)
	}
	
	return dto.GetURLOutput{RedirectURL: rawUrl}, nil
}

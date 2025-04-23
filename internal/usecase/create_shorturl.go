package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/erlitx/link_shortner/internal/domain"
	"github.com/erlitx/link_shortner/internal/dto"
)

func (u *UseCase) CreateShortURL(ctx context.Context, input dto.CreateShortUrlInput) (dto.CreateShortUrlOutput, error) {
	var output dto.CreateShortUrlOutput

	shortUrl := GenerateShortIDFromURL(input.RawURL, 8)

	url, err := domain.NewUrl(input.RawURL, shortUrl)

	u.cache.Set(url)
	
	if err != nil {
		return output, fmt.Errorf("domain.NewUrl: %w", err)
	}



	err = u.postgres.CreateShortURL(ctx, url)
	if err != nil {
		return output, fmt.Errorf("u.postgres.CreateURL: %w", err)
	}
	output = dto.CreateShortUrlOutput{input.Host + "/" + string(url.ShortURL)}
	return output, nil
}


func GenerateShortIDFromURL(input string, length int) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	sum := hasher.Sum(nil) // 20 bytes

	encoded := base64.URLEncoding.EncodeToString(sum)

	if len(encoded) > length {
		encoded = encoded[:length]
	}

	return encoded
}

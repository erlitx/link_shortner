package usecase

import (
	"context"
	"fmt"
	"path"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/skip2/go-qrcode"
)

func (u *UseCase) GenerateQRCode(ctx context.Context, m kafka.Message) error {
	// Some logic to generate QR code from kafka message value and save it to S3
	log.Info().
		Str("topic", m.Topic).
		Int("partition", m.Partition).
		Int64("offset", m.Offset).
		Str("key", string(m.Key)).
		Str("value", string(m.Value)).
		Msg("URL recieved")

	url := "http://" + string(m.Value)

	// Генерация QR
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate QR")
		return err
	}

	filename := path.Base(string(m.Value))

	key := fmt.Sprintf("%s.png", filename)

	// Загрузка в S3
	err = u.storage.SaveFile(ctx, key, png)
	if err != nil {
		log.Error().Err(err).Msg("failed to upload QR to MinIO")
		return err
	}
	log.Info().Str("s3_key", key).Msg("QR code saved to MinIO")
	return nil
}

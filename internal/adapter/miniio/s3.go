package minioadapter

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Host     string `envconfig:"MINIIO_HOST"       required:"true"`
	User     string `envconfig:"MINIIO_USER"     required:"true"`
	Password string `envconfig:"MINIIO_PASSWORD" required:"true"`
	Bucket   string `envconfig:"MINIIO_BUCKET" required:"true"`
}

type MinioClient struct {
	client *minio.Client
	bucket string
}

func New(c Config) (*MinioClient, error) {
	// Determine if HTTPS should be used
	useSSL := false
	if c.Host[:5] == "https" {
		useSSL = true
	}



	// Initialize MinIO client
	minioClient, err := minio.New(c.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(c.User, c.Password, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %w", err)
	}


	return &MinioClient{
		client: minioClient,
		bucket: c.Bucket,
	}, nil
}

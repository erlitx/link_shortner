package minioadapter

import (
	"bytes"
	"context"
	"fmt"
	"mime"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)


func (m *MinioClient) SaveFile(ctx context.Context, key string, data []byte) error {
	fmt.Println("---------- FILE SAVE TO MINIO ----------")

	// Detect content type from file extension
	contentType := mime.TypeByExtension(filepath.Ext(key))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Prepare reader and size
	reader := bytes.NewReader(data)
	size := int64(len(data))

	// Upload the object
	_, err := m.client.PutObject(ctx, m.bucket, key, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("failed to upload to MinIO: %w", err)
	}

	fmt.Println("✅ File uploaded:", key)
	return nil
}

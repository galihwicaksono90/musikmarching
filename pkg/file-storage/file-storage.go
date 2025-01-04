package fileStorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"

	"galihwicaksono90/musikmarching-be/pkg/config"
)

func NewStorage(logger *logrus.Logger, config config.Config) *minio.Client {
	// minioClient, err := minio.New(config.MinioAddress, &minio.Options{
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: config.MinioSecure,
	})

	if err != nil {
		logger.Fatalln(err)
	}

	return minioClient
}

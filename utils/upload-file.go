package utils

import (
	"context"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
)

func UploadFile(fileStorage *minio.Client, file multipart.File, header *multipart.FileHeader, location model.FileLocation) (minio.UploadInfo, error) {
  bucketName := viper.GetString("MINIO_BUCKET_NAME")
  ctx := context.Background()

  id := uuid.New()
  timestamp := time.Now().Format("2006-01-02")
  fileName := fmt.Sprintf("%s/%s/%s/%s", location, timestamp, id, header.Filename)

  options := minio.PutObjectOptions{
    ContentType: header.Header.Get("Content-Type"),
  }

  return fileStorage.PutObject(
    ctx, 
    bucketName, 
    fileName,
    file, 
    -1, 
    options,
  )
}

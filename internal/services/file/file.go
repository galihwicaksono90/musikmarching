package file

import (
	"bytes"
	"context"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	"image/png"
	"io"
	"path/filepath"
	"strings"

	"mime/multipart"
	"time"

	"net/http"

	"github.com/gen2brain/go-fitz"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type FileService interface {
	UploadPdfFile(*http.Request, string) (url string, images []string, err error)
	UploadAudioFile(*http.Request, string) (url string, err error)
}

type fileService struct {
	logger      *logrus.Logger
	fileStorage *minio.Client
}

// UploadPdfFile implements ScoreService.
func (s *fileService) UploadPdfFile(r *http.Request, name string) (string, []string, error) {
	file, header, err := r.FormFile(name)
	if err != nil || file == nil {
		s.logger.Errorln(err)
		return "", []string{}, err
	}
	defer file.Close()

	pdfUploadInfo, err := uploadFile(s.fileStorage, file, header, model.PDF_LOCATION)
	if err != nil {
		s.logger.Errorln(err)
		return "", []string{}, err
	}

	pages := 2
	images, err := s.pdfToImages(file, header.Filename, pages)
	if err != nil {
		s.logger.Errorln(err)
		return "", []string{}, err
	}

	return pdfUploadInfo.Location, images, nil
}

func (s *fileService) pdfToImages(file multipart.File, fileName string, pages int) ([]string, error) {
	images := []string{}
	seeker, ok := file.(io.Seeker)
	if !ok {
		return images, nil
	}
	_, err := seeker.Seek(0, io.SeekStart)
	if err != nil {
		return images, err
	}

	doc, err := fitz.NewFromReader(file)
	if err != nil {
		return images, err
	}
	defer doc.Close()

	pageCount := doc.NumPage()
	if pageCount == 0 {
		return images, err
	}

	if pageCount < pages {
		pages = pageCount
	}

	id := uuid.New()
	timestamp := time.Now().Format("2006-01-02")
	fileName = fmt.Sprintf("%s/%s/%s/%s", model.PDF_IMAGE_LOCATION, timestamp, id, strings.TrimSuffix(fileName, filepath.Ext(fileName)))

	for i := 0; i < pages; i++ {
		img, err := doc.Image(i)
		if err != nil {
			return images, fmt.Errorf("Failed to extract page %d %v", i+1, err)
		}

		imgBuffer := new(bytes.Buffer)

		if err := png.Encode(imgBuffer, img); err != nil {
			return images, fmt.Errorf("Failed to encode image %v", err)
		}

		bucketName := viper.GetString("MINIO_BUCKET_NAME")
		result, err := s.fileStorage.PutObject(
			context.Background(),
			bucketName,
			fmt.Sprintf("%s-page%d.png", fileName, i+1),
			imgBuffer,
			-1,
			minio.PutObjectOptions{
				ContentType: "image/png",
			},
		)
		if err != nil {
			return images, fmt.Errorf("Failed to upload image %v", err)
		}

		images = append(images, result.Location)
	}

	return images, nil
}

// UploadAudioFile implements ScoreService.
func (s *fileService) UploadAudioFile(r *http.Request, name string) (url string, err error) {
	file, header, err := r.FormFile(name)
	if err != nil {
		s.logger.Errorln(err)
		return "", err
	}
	defer file.Close()

	pdfUploadInfo, err := uploadFile(s.fileStorage, file, header, model.AUDIO_LOCATION)
	if err != nil {
		s.logger.Errorln(err)
		return "", err
	}

	return pdfUploadInfo.Location, nil
}

func uploadFile(fileStorage *minio.Client, file multipart.File, header *multipart.FileHeader, location model.FileLocation) (minio.UploadInfo, error) {
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

func NewFileService(logger *logrus.Logger, fileStorage *minio.Client) FileService {
	return &fileService{
		logger,
		fileStorage,
	}
}

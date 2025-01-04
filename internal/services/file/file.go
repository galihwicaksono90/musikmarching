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
	UploadPdfFile(*http.Request, string, int) (string, []string, error)
	UploadAudioFile(*http.Request, string) (string, error)
	UploadPaymentProof(*http.Request, string) (string, error)
}

type fileService struct {
	logger      *logrus.Logger
	fileStorage *minio.Client
}

// UploadPdfFile implements ScoreService.
func (s *fileService) UploadPdfFile(r *http.Request, name string, pages int) (string, []string, error) {
	file, header, err := r.FormFile(name)
	if err != nil || file == nil {
		s.logger.Errorln(err)
		return "", []string{}, err
	}
	defer file.Close()

	fileName := s.generateFileName(model.PDF_LOCATION, header.Filename)
	pdfUploadInfo, err := s.uploadFile(file, fileName, header.Header.Get("Content-Type"))
	if err != nil {
		s.logger.Errorln(err)
		return "", []string{}, err
	}

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

		fileName := fmt.Sprintf("%s-page%d.png", fileName, i+1)
		res, err := s.uploadFile(imgBuffer, fileName, "image/png")
		if err != nil {
			return images, fmt.Errorf("Failed to upload image %v", err)
		}

		images = append(images, res.Location)
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

	fileName := s.generateFileName(model.AUDIO_LOCATION, header.Filename)
	pdfUploadInfo, err := s.uploadFile(file, fileName, header.Header.Get("Content-Type"))
	if err != nil {
		s.logger.Errorln(err)
		return "", err
	}

	return pdfUploadInfo.Location, nil
}

func (s *fileService) UploadPaymentProof(r *http.Request, name string) (string, error) {
	file, header, err := r.FormFile(name)
	if err != nil {
		s.logger.Errorln(err)
		return "", err
	}
	defer file.Close()

	fileType := header.Header.Get("Content-Type")
	fileName := s.generateFileName(model.PAYMENT_PROOF_IMAGE_LOCATION, header.Filename)

	result, err := s.uploadFile(file, fileName, fileType)

	return result.Location, err
}

func (s *fileService) uploadFile(file io.Reader, fileName string, fileType string) (minio.UploadInfo, error) {
	bucketName := viper.GetString("MINIO_BUCKET_NAME")
	s.logger.Infoln("bucketName", bucketName)
	s.logger.Infoln("fileName", fileName)
	s.logger.Infoln("fileType", fileType)

	return s.fileStorage.PutObject(
		context.Background(),
		bucketName,
		fileName,
		file,
		-1,
		minio.PutObjectOptions{
			ContentType: fileType,
		},
	)
}

func (s *fileService) generateFileName(location model.FileLocation, fileName string) string {
	id := uuid.New()
	timestamp := time.Now().Format("2006-01-02")
	name := fmt.Sprintf("%s/%s/%s/%s", location, timestamp, id, fileName)
	return name
}

func NewFileService(logger *logrus.Logger, fileStorage *minio.Client) FileService {
	return &fileService{
		logger,
		fileStorage,
	}
}

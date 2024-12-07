package validator

import (
	"mime"
	"mime/multipart"
	"path/filepath"

	validator "github.com/go-playground/validator/v10"
)

type FileUploadRequest struct {
	File       *multipart.File       `validate:"required"`
	FileHeader *multipart.FileHeader `validate:"required"`
}

func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("pdf", validatePdfFileType)
	validate.RegisterValidation("mp3", validateMp3FileType)

	return validate
}

func validatePdfFileType(fl validator.FieldLevel) bool {
	return validateFileType(fl, "application/pdf")
}

func validateMp3FileType(fl validator.FieldLevel) bool {
	return validateFileType(fl, "audio/mpeg")
}

func validateFileType(fl validator.FieldLevel, fileType string) bool {
	file, ok := fl.Parent().Interface().(FileUploadRequest)
	if !ok {
		return false
	}

	ext := filepath.Ext(file.FileHeader.Filename)
	mimeType := mime.TypeByExtension(ext)

	return mimeType == fileType
}

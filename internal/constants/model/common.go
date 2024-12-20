package model

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"time"

	"github.com/google/uuid"
)

type SessionUser struct {
	ID         uuid.UUID   `json:"id"`
	Email      string      `json:"email"`
	Name       string      `json:"name"`
	RoleName   db.Rolename `json:"role_name"`
	PictureUrl string      `json:"picture"`
}

type Account struct {
	ID        uuid.UUID   `json:"id"`
	Email     string      `json:"email"`
	Name      string      `json:"name"`
	RoleName  db.Rolename `json:"role_name"`
	Picture   string      `json:"picture"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt time.Time   `json:"deleted_at"`
}

type FileLocation string

const (
    PDF_LOCATION FileLocation = "pdf"
    AUDIO_LOCATION  FileLocation = "audio"
    PDF_IMAGE_LOCATION  FileLocation = "pdf_image"
)

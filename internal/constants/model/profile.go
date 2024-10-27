package model

import (
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
)

type Profile struct {
	*db.Profile
	UploadedScores *[]db.Score `json:"uploaded_scores"`
}


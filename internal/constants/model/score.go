package model

import (
	"math/big"

	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateScoreDTO struct {
	ContributorID uuid.UUID      `json:"contributor_id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Price         *big.Int       `json:"price"`
	PdfUrl        string         `json:"pdf_url"`
	PdfImageUrls  []string       `json:"pdf_image_urls"`
	AudioUrl      string         `json:"audio_url"`
	Difficulty    db.Difficulty  `json:"difficulty"`
	ContentType   db.ContentType `json:"content_type"`
}

type UpdateScoreDTO struct {
	ContributorID uuid.UUID          `json:"contributor_id"`
	Title         pgtype.Text        `json:"title"`
	Description   pgtype.Text        `json:"description"`
	Difficulty    db.NullDifficulty  `json:"difficulty"`
	ContentType   db.NullContentType `json:"content_type"`
	Price         pgtype.Numeric     `json:"price"`
	PdfUrl        pgtype.Text        `json:"pdf_url"`
	PdfImageUrls  []string           `json:"pdf_image_urls"`
	AudioUrl      pgtype.Text        `json:"audio_url"`
}

type GetAllPublicScoreDTO struct {
	Limit       int32
	Offset      int32
	Title       string
	Instruments []string
	Categories  []string
	Allocations []string
}

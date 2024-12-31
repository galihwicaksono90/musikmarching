package model

import (
	"math/big"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateScoreDTO struct {
	ContributorID uuid.UUID `json:"contributor_id"`
	Title         string    `json:"title"`
	Price         *big.Int  `json:"price"`
	PdfUrl        string    `json:"pdf_url"`
	PdfImageUrls  []string  `json:"pdf_image_urls"`
	AudioUrl      string    `json:"audio_url"`
}

type UpdateScoreDTO struct {
	ContributorID uuid.UUID      `json:"contributor_id"`
	Title         pgtype.Text    `db:"title" json:"title"`
	Price         pgtype.Numeric `db:"price" json:"price"`
	PdfUrl        pgtype.Text    `json:"pdf_url"`
	PdfImageUrls  []string       `json:"pdf_image_urls"`
	AudioUrl      pgtype.Text    `json:"audio_url"`
}

type GetAllPublicScoreDTO struct {
	Limit    int32
	Offset   int32
	Title    string
	Instruments []string
	Categories []string
	Allocations []string
}


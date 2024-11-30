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
	MusicUrl      string    `json:"music_url"`
}

type UpdateScoreDTO struct {
	ContributorID uuid.UUID   `json:"contributor_id"`
	Title pgtype.Text    `db:"title" json:"title"`
	Price pgtype.Numeric `db:"price" json:"price"`
	// Price         *big.Int   `json:"price"`
	// PdfUrl        string   `json:"pdf_url"`
	// MusicUrl      string   `json:"music_url"`
}

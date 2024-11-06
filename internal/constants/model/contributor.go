package model

import "time"

type Contributor struct {
	ID         string    `json:"id"`
	IsVerified bool      `json:"isVerified"`
	VerifiedAt time.Time `json:"verifiedAt"`
}

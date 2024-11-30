// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error)
	CreateContributor(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	CreatePurchase(ctx context.Context, arg CreatePurchaseParams) (uuid.UUID, error)
	CreateScore(ctx context.Context, arg CreateScoreParams) (uuid.UUID, error)
	GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error)
	GetAccountById(ctx context.Context, id uuid.UUID) (GetAccountByIdRow, error)
	GetAccounts(ctx context.Context) ([]GetAccountsRow, error)
	GetContributorById(ctx context.Context, id uuid.UUID) (GetContributorByIdRow, error)
	GetPurchaseByAccountAndScoreId(ctx context.Context, arg GetPurchaseByAccountAndScoreIdParams) (Purchase, error)
	GetPurchaseById(ctx context.Context, id uuid.UUID) (Purchase, error)
	GetPurchases(ctx context.Context, accountID uuid.UUID) ([]Purchase, error)
	GetScoreByContributorId(ctx context.Context, id uuid.UUID) ([]Score, error)
	GetScoreById(ctx context.Context, id uuid.UUID) (Score, error)
	GetUnverifiedContributors(ctx context.Context) ([]Contributor, error)
	GetVerifiedScoreById(ctx context.Context, id uuid.UUID) (GetVerifiedScoreByIdRow, error)
	GetVerifiedScores(ctx context.Context, arg GetVerifiedScoresParams) ([]GetVerifiedScoresRow, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (uuid.UUID, error)
	UpdateAccountRole(ctx context.Context, arg UpdateAccountRoleParams) (uuid.UUID, error)
	UpdateScore(ctx context.Context, arg UpdateScoreParams) error
	VerifyContributor(ctx context.Context, id uuid.UUID) error
}

var _ Querier = (*Queries)(nil)

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
	CreateAllocation(ctx context.Context, name string) (Allocation, error)
	CreateCategory(ctx context.Context, name string) (Category, error)
	CreateContributor(ctx context.Context, arg CreateContributorParams) (uuid.UUID, error)
	CreateInstrument(ctx context.Context, name string) (Instrument, error)
	CreatePurchase(ctx context.Context, arg CreatePurchaseParams) (uuid.UUID, error)
	CreateScore(ctx context.Context, arg CreateScoreParams) (uuid.UUID, error)
	DeleteAllocation(ctx context.Context, id int32) error
	DeleteCategory(ctx context.Context, id int32) error
	DeleteInstrument(ctx context.Context, id int32) error
	DeleteScoreAllocation(ctx context.Context, arg DeleteScoreAllocationParams) error
	DeleteScoreCategory(ctx context.Context, arg DeleteScoreCategoryParams) error
	DeleteScoreInstrument(ctx context.Context, arg DeleteScoreInstrumentParams) error
	GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error)
	GetAccountById(ctx context.Context, id uuid.UUID) (GetAccountByIdRow, error)
	GetAccounts(ctx context.Context) ([]GetAccountsRow, error)
	GetAllContributors(ctx context.Context) ([]ContributorAccountScore, error)
	GetAllPublicScores(ctx context.Context, arg GetAllPublicScoresParams) ([]ScorePublicView, error)
	GetAllPurchases(ctx context.Context) ([]Purchase, error)
	GetAllocations(ctx context.Context) ([]Allocation, error)
	GetCategories(ctx context.Context) ([]Category, error)
	GetContributorById(ctx context.Context, id uuid.UUID) (ContributorAccountScore, error)
	GetInstruments(ctx context.Context) ([]Instrument, error)
	GetPurchaseByAccountAndScoreId(ctx context.Context, arg GetPurchaseByAccountAndScoreIdParams) (Purchase, error)
	GetPurchaseById(ctx context.Context, arg GetPurchaseByIdParams) (Purchase, error)
	GetPurchasesByAccountId(ctx context.Context, accountID uuid.UUID) ([]Purchase, error)
	GetScoreByContributorID(ctx context.Context, arg GetScoreByContributorIDParams) (GetScoreByContributorIDRow, error)
	GetScoreByContributorId(ctx context.Context, id uuid.UUID) ([]Score, error)
	GetScoreById(ctx context.Context, id uuid.UUID) (Score, error)
	GetScores(ctx context.Context, arg GetScoresParams) ([]GetScoresRow, error)
	GetScoresByContributorID(ctx context.Context, arg GetScoresByContributorIDParams) ([]GetScoresByContributorIDRow, error)
	GetScoresPaginated(ctx context.Context) ([]Score, error)
	GetUnverifiedContributors(ctx context.Context) ([]Contributor, error)
	GetVerifiedScoreById(ctx context.Context, id uuid.UUID) (GetVerifiedScoreByIdRow, error)
	GetVerifiedScores(ctx context.Context, arg GetVerifiedScoresParams) ([]GetVerifiedScoresRow, error)
	InsertScoreAllocation(ctx context.Context, arg InsertScoreAllocationParams) error
	InsertScoreCategory(ctx context.Context, arg InsertScoreCategoryParams) error
	InsertScoreInstrument(ctx context.Context, arg InsertScoreInstrumentParams) error
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (uuid.UUID, error)
	UpdateAccountRole(ctx context.Context, arg UpdateAccountRoleParams) (uuid.UUID, error)
	UpdateAllocation(ctx context.Context, arg UpdateAllocationParams) error
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error
	UpdateInstrument(ctx context.Context, arg UpdateInstrumentParams) error
	UpdatePurchaseProof(ctx context.Context, arg UpdatePurchaseProofParams) error
	UpdateScore(ctx context.Context, arg UpdateScoreParams) error
	VerifyContributor(ctx context.Context, id uuid.UUID) error
	VerifyPurchase(ctx context.Context, id uuid.UUID) error
	VerifyScore(ctx context.Context, id uuid.UUID) error
}

var _ Querier = (*Queries)(nil)

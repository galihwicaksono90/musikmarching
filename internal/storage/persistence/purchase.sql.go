// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: purchase.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createPurchase = `-- name: CreatePurchase :one
insert into purchase (account_id, score_id, price, title)
values ($1, $2, $3, $4)
returning id
`

type CreatePurchaseParams struct {
	AccountID uuid.UUID      `db:"account_id" json:"account_id"`
	ScoreID   uuid.UUID      `db:"score_id" json:"score_id"`
	Price     pgtype.Numeric `db:"price" json:"price"`
	Title     string         `db:"title" json:"title"`
}

func (q *Queries) CreatePurchase(ctx context.Context, arg CreatePurchaseParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createPurchase,
		arg.AccountID,
		arg.ScoreID,
		arg.Price,
		arg.Title,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getPurchaseByAccountAndScoreId = `-- name: GetPurchaseByAccountAndScoreId :one
select id, invoice_serial, account_id, score_id, price, title, is_verified, verifiedat, created_at, updated_at, deleted_at from purchase
where account_id = $1 and score_id = $2
`

type GetPurchaseByAccountAndScoreIdParams struct {
	AccountID uuid.UUID `db:"account_id" json:"account_id"`
	ScoreID   uuid.UUID `db:"score_id" json:"score_id"`
}

func (q *Queries) GetPurchaseByAccountAndScoreId(ctx context.Context, arg GetPurchaseByAccountAndScoreIdParams) (Purchase, error) {
	row := q.db.QueryRow(ctx, getPurchaseByAccountAndScoreId, arg.AccountID, arg.ScoreID)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.InvoiceSerial,
		&i.AccountID,
		&i.ScoreID,
		&i.Price,
		&i.Title,
		&i.IsVerified,
		&i.Verifiedat,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getPurchaseById = `-- name: GetPurchaseById :one
select id, invoice_serial, account_id, score_id, price, title, is_verified, verifiedat, created_at, updated_at, deleted_at from purchase
where id = $1
`

func (q *Queries) GetPurchaseById(ctx context.Context, id uuid.UUID) (Purchase, error) {
	row := q.db.QueryRow(ctx, getPurchaseById, id)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.InvoiceSerial,
		&i.AccountID,
		&i.ScoreID,
		&i.Price,
		&i.Title,
		&i.IsVerified,
		&i.Verifiedat,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getPurchases = `-- name: GetPurchases :many
select id, invoice_serial, account_id, score_id, price, title, is_verified, verifiedat, created_at, updated_at, deleted_at from purchase
where account_id = $1
order by created_at desc
`

func (q *Queries) GetPurchases(ctx context.Context, accountID uuid.UUID) ([]Purchase, error) {
	rows, err := q.db.Query(ctx, getPurchases, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Purchase{}
	for rows.Next() {
		var i Purchase
		if err := rows.Scan(
			&i.ID,
			&i.InvoiceSerial,
			&i.AccountID,
			&i.ScoreID,
			&i.Price,
			&i.Title,
			&i.IsVerified,
			&i.Verifiedat,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

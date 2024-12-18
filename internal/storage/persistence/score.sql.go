// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: score.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createScore = `-- name: CreateScore :one
insert into score (
  title,
  price,
  pdf_url,
  music_url,
  contributor_id
) values (
  $1,
  $2,
  $3,
  $4,
  $5
) returning id
`

type CreateScoreParams struct {
	Title         string         `db:"title" json:"title"`
	Price         pgtype.Numeric `db:"price" json:"price"`
	PdfUrl        pgtype.Text    `db:"pdf_url" json:"pdf_url"`
	MusicUrl      pgtype.Text    `db:"music_url" json:"music_url"`
	ContributorID uuid.UUID      `db:"contributor_id" json:"contributor_id"`
}

func (q *Queries) CreateScore(ctx context.Context, arg CreateScoreParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createScore,
		arg.Title,
		arg.Price,
		arg.PdfUrl,
		arg.MusicUrl,
		arg.ContributorID,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getScoreByContributorId = `-- name: GetScoreByContributorId :many
select id, contributor_id, title, price, is_verified, verified_at, pdf_url, music_url, created_at, updated_at, deleted_at
from score
where contributor_id = $1
`

func (q *Queries) GetScoreByContributorId(ctx context.Context, id uuid.UUID) ([]Score, error) {
	rows, err := q.db.Query(ctx, getScoreByContributorId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Score{}
	for rows.Next() {
		var i Score
		if err := rows.Scan(
			&i.ID,
			&i.ContributorID,
			&i.Title,
			&i.Price,
			&i.IsVerified,
			&i.VerifiedAt,
			&i.PdfUrl,
			&i.MusicUrl,
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

const getScoreById = `-- name: GetScoreById :one
select id, contributor_id, title, price, is_verified, verified_at, pdf_url, music_url, created_at, updated_at, deleted_at
from score s
where s.id = $1
`

func (q *Queries) GetScoreById(ctx context.Context, id uuid.UUID) (Score, error) {
	row := q.db.QueryRow(ctx, getScoreById, id)
	var i Score
	err := row.Scan(
		&i.ID,
		&i.ContributorID,
		&i.Title,
		&i.Price,
		&i.IsVerified,
		&i.VerifiedAt,
		&i.PdfUrl,
		&i.MusicUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getScores = `-- name: GetScores :many
select s.id, s.title, s.is_verified, s.price, a.name, a.email
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
order by s.created_at desc
limit $2::int
offset $1::int
`

type GetScoresParams struct {
	Pageoffset int32 `db:"pageoffset" json:"pageoffset"`
	Pagelimit  int32 `db:"pagelimit" json:"pagelimit"`
}

type GetScoresRow struct {
	ID         uuid.UUID      `db:"id" json:"id"`
	Title      string         `db:"title" json:"title"`
	IsVerified bool           `db:"is_verified" json:"is_verified"`
	Price      pgtype.Numeric `db:"price" json:"price"`
	Name       string         `db:"name" json:"name"`
	Email      string         `db:"email" json:"email"`
}

func (q *Queries) GetScores(ctx context.Context, arg GetScoresParams) ([]GetScoresRow, error) {
	rows, err := q.db.Query(ctx, getScores, arg.Pageoffset, arg.Pagelimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetScoresRow{}
	for rows.Next() {
		var i GetScoresRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.IsVerified,
			&i.Price,
			&i.Name,
			&i.Email,
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

const getScoresByContributorID = `-- name: GetScoresByContributorID :many
select s.id, s.title, s.is_verified, s.price, a.name, a.email
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
where s.contributor_id = $1
order by s.is_verified desc, s.created_at desc
limit $3::int
offset $2::int
`

type GetScoresByContributorIDParams struct {
	ID         uuid.UUID `db:"id" json:"id"`
	Pageoffset int32     `db:"pageoffset" json:"pageoffset"`
	Pagelimit  int32     `db:"pagelimit" json:"pagelimit"`
}

type GetScoresByContributorIDRow struct {
	ID         uuid.UUID      `db:"id" json:"id"`
	Title      string         `db:"title" json:"title"`
	IsVerified bool           `db:"is_verified" json:"is_verified"`
	Price      pgtype.Numeric `db:"price" json:"price"`
	Name       string         `db:"name" json:"name"`
	Email      string         `db:"email" json:"email"`
}

func (q *Queries) GetScoresByContributorID(ctx context.Context, arg GetScoresByContributorIDParams) ([]GetScoresByContributorIDRow, error) {
	rows, err := q.db.Query(ctx, getScoresByContributorID, arg.ID, arg.Pageoffset, arg.Pagelimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetScoresByContributorIDRow{}
	for rows.Next() {
		var i GetScoresByContributorIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.IsVerified,
			&i.Price,
			&i.Name,
			&i.Email,
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

const getScoresPaginated = `-- name: GetScoresPaginated :many
select id, contributor_id, title, price, is_verified, verified_at, pdf_url, music_url, created_at, updated_at, deleted_at
from score
where deleted_at is null
order by
    case when $3 = 'price_asc' then price when $3 = 'price_desc' then price end,
    case
        when $3 = 'created_at_asc'
        then created_at
        when $3 = 'created_at_desc'
        then created_at
    end desc
limit $1
offset $2
`

type GetScoresPaginatedParams struct {
	Limit   int32       `db:"limit" json:"limit"`
	Offset  int32       `db:"offset" json:"offset"`
	Column3 interface{} `db:"column_3" json:"column_3"`
}

func (q *Queries) GetScoresPaginated(ctx context.Context, arg GetScoresPaginatedParams) ([]Score, error) {
	rows, err := q.db.Query(ctx, getScoresPaginated, arg.Limit, arg.Offset, arg.Column3)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Score{}
	for rows.Next() {
		var i Score
		if err := rows.Scan(
			&i.ID,
			&i.ContributorID,
			&i.Title,
			&i.Price,
			&i.IsVerified,
			&i.VerifiedAt,
			&i.PdfUrl,
			&i.MusicUrl,
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

const getVerifiedScoreById = `-- name: GetVerifiedScoreById :one
select s.id, s.title, s.price, a.name as contributor_name
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
where s.is_verified = true and s.id = $1
`

type GetVerifiedScoreByIdRow struct {
	ID              uuid.UUID      `db:"id" json:"id"`
	Title           string         `db:"title" json:"title"`
	Price           pgtype.Numeric `db:"price" json:"price"`
	ContributorName string         `db:"contributor_name" json:"contributor_name"`
}

func (q *Queries) GetVerifiedScoreById(ctx context.Context, id uuid.UUID) (GetVerifiedScoreByIdRow, error) {
	row := q.db.QueryRow(ctx, getVerifiedScoreById, id)
	var i GetVerifiedScoreByIdRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Price,
		&i.ContributorName,
	)
	return i, err
}

const getVerifiedScores = `-- name: GetVerifiedScores :many
select s.id, s.title, s.price, a.name, a.email
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
where s.is_verified = true and c.is_verified = true
limit $2::int
offset $1::int
`

type GetVerifiedScoresParams struct {
	Pageoffset int32 `db:"pageoffset" json:"pageoffset"`
	Pagelimit  int32 `db:"pagelimit" json:"pagelimit"`
}

type GetVerifiedScoresRow struct {
	ID    uuid.UUID      `db:"id" json:"id"`
	Title string         `db:"title" json:"title"`
	Price pgtype.Numeric `db:"price" json:"price"`
	Name  string         `db:"name" json:"name"`
	Email string         `db:"email" json:"email"`
}

func (q *Queries) GetVerifiedScores(ctx context.Context, arg GetVerifiedScoresParams) ([]GetVerifiedScoresRow, error) {
	rows, err := q.db.Query(ctx, getVerifiedScores, arg.Pageoffset, arg.Pagelimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetVerifiedScoresRow{}
	for rows.Next() {
		var i GetVerifiedScoresRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Price,
			&i.Name,
			&i.Email,
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

const updateScore = `-- name: UpdateScore :exec
update score set
  title = COALESCE($1, title),
  price = COALESCE($2, price),
  updated_at = now()
where id = $3
`

type UpdateScoreParams struct {
	Title pgtype.Text    `db:"title" json:"title"`
	Price pgtype.Numeric `db:"price" json:"price"`
	ID    uuid.UUID      `db:"id" json:"id"`
}

func (q *Queries) UpdateScore(ctx context.Context, arg UpdateScoreParams) error {
	_, err := q.db.Exec(ctx, updateScore, arg.Title, arg.Price, arg.ID)
	return err
}

const verifyScore = `-- name: VerifyScore :exec
update score set
  is_verified = true,
  verified_at = now()
where id = $1
`

func (q *Queries) VerifyScore(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, verifyScore, id)
	return err
}

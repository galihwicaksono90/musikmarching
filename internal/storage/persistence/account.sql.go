// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: account.sql

package persistence

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO Account (email, name, picture_url, role_id)
VALUES ($1, $2, $3, (select id from role where name = 'user'))
RETURNING id, name
`

type CreateAccountParams struct {
	Email      string      `db:"email" json:"email"`
	Name       string      `db:"name" json:"name"`
	Pictureurl pgtype.Text `db:"pictureurl" json:"pictureurl"`
}

type CreateAccountRow struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Email, arg.Name, arg.Pictureurl)
	var i CreateAccountRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
select a.id, a.name, a.email, a.picture_url, a.created_at, a.updated_at, a.deleted_at, a.role_id, r.name as role_name from account a
inner join role r on r.id = a.role_id
where a.email = $1
limit 1
`

type GetAccountByEmailRow struct {
	ID         uuid.UUID          `db:"id" json:"id"`
	Name       string             `db:"name" json:"name"`
	Email      string             `db:"email" json:"email"`
	PictureUrl pgtype.Text        `db:"picture_url" json:"picture_url"`
	CreatedAt  time.Time          `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamptz `db:"deleted_at" json:"deleted_at"`
	RoleID     uuid.UUID          `db:"role_id" json:"role_id"`
	RoleName   Rolename           `db:"role_name" json:"role_name"`
}

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error) {
	row := q.db.QueryRow(ctx, getAccountByEmail, email)
	var i GetAccountByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PictureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.RoleID,
		&i.RoleName,
	)
	return i, err
}

const getAccountById = `-- name: GetAccountById :one
select a.id, a.name, a.email, a.picture_url, a.created_at, a.updated_at, a.deleted_at, a.role_id, r.name as role_name from account a
inner join role r on r.id = a.role_id
where a.id = $1
limit 1
`

type GetAccountByIdRow struct {
	ID         uuid.UUID          `db:"id" json:"id"`
	Name       string             `db:"name" json:"name"`
	Email      string             `db:"email" json:"email"`
	PictureUrl pgtype.Text        `db:"picture_url" json:"picture_url"`
	CreatedAt  time.Time          `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamptz `db:"deleted_at" json:"deleted_at"`
	RoleID     uuid.UUID          `db:"role_id" json:"role_id"`
	RoleName   Rolename           `db:"role_name" json:"role_name"`
}

func (q *Queries) GetAccountById(ctx context.Context, id uuid.UUID) (GetAccountByIdRow, error) {
	row := q.db.QueryRow(ctx, getAccountById, id)
	var i GetAccountByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PictureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.RoleID,
		&i.RoleName,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :many
select a.id, a.name, a.email, a.picture_url, a.created_at, a.updated_at, a.deleted_at, a.role_id, r.name as role_name from account a
inner join role r on r.id = a.role_id
`

type GetAccountsRow struct {
	ID         uuid.UUID          `db:"id" json:"id"`
	Name       string             `db:"name" json:"name"`
	Email      string             `db:"email" json:"email"`
	PictureUrl pgtype.Text        `db:"picture_url" json:"picture_url"`
	CreatedAt  time.Time          `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamptz `db:"deleted_at" json:"deleted_at"`
	RoleID     uuid.UUID          `db:"role_id" json:"role_id"`
	RoleName   Rolename           `db:"role_name" json:"role_name"`
}

func (q *Queries) GetAccounts(ctx context.Context) ([]GetAccountsRow, error) {
	rows, err := q.db.Query(ctx, getAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsRow{}
	for rows.Next() {
		var i GetAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.PictureUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.RoleID,
			&i.RoleName,
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

const updateAccount = `-- name: UpdateAccount :one
update account a
set name = $1, picture_url = $2
where id = $3
returning a.id
`

type UpdateAccountParams struct {
	Name       string      `db:"name" json:"name"`
	Pictureurl pgtype.Text `db:"pictureurl" json:"pictureurl"`
	ID         uuid.UUID   `db:"id" json:"id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, updateAccount, arg.Name, arg.Pictureurl, arg.ID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const updateAccountRole = `-- name: UpdateAccountRole :one
update account a
set role_id = (select id from role as r where r.name = $1)
where a.id = $2
RETURNING id
`

type UpdateAccountRoleParams struct {
	Rolename Rolename  `db:"rolename" json:"rolename"`
	ID       uuid.UUID `db:"id" json:"id"`
}

func (q *Queries) UpdateAccountRole(ctx context.Context, arg UpdateAccountRoleParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, updateAccountRole, arg.Rolename, arg.ID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

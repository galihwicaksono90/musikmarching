// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: instrument.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createInstrument = `-- name: CreateInstrument :one
insert into instrument (name) values ($1) returning id, name
`

func (q *Queries) CreateInstrument(ctx context.Context, name string) (Instrument, error) {
	row := q.db.QueryRow(ctx, createInstrument, name)
	var i Instrument
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createScoreInstrument = `-- name: CreateScoreInstrument :exec
insert into score_instrument (score_id, instrument_id) values ($1, $2)
`

type CreateScoreInstrumentParams struct {
	ScoreID      uuid.UUID `db:"score_id" json:"score_id"`
	InstrumentID int32     `db:"instrument_id" json:"instrument_id"`
}

func (q *Queries) CreateScoreInstrument(ctx context.Context, arg CreateScoreInstrumentParams) error {
	_, err := q.db.Exec(ctx, createScoreInstrument, arg.ScoreID, arg.InstrumentID)
	return err
}

const deleteInstrument = `-- name: DeleteInstrument :exec
delete from instrument where id = $1
`

func (q *Queries) DeleteInstrument(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteInstrument, id)
	return err
}

const deleteScoreInstrument = `-- name: DeleteScoreInstrument :exec
delete from score_instrument where score_id = $1
`

func (q *Queries) DeleteScoreInstrument(ctx context.Context, scoreID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteScoreInstrument, scoreID)
	return err
}

const getInstruments = `-- name: GetInstruments :many
select id, name from instrument
`

func (q *Queries) GetInstruments(ctx context.Context) ([]Instrument, error) {
	rows, err := q.db.Query(ctx, getInstruments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Instrument{}
	for rows.Next() {
		var i Instrument
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInstrument = `-- name: UpdateInstrument :exec
update instrument set name = $1 where id = $2
`

type UpdateInstrumentParams struct {
	Name string `db:"name" json:"name"`
	ID   int32  `db:"id" json:"id"`
}

func (q *Queries) UpdateInstrument(ctx context.Context, arg UpdateInstrumentParams) error {
	_, err := q.db.Exec(ctx, updateInstrument, arg.Name, arg.ID)
	return err
}

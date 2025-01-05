-- name: GetInstruments :many
select * from instrument;

-- name: DeleteInstrument :exec
delete from instrument where id = $1;

-- name: CreateInstrument :one
insert into instrument (name) values ($1) returning *;

-- name: UpdateInstrument :exec
update instrument set name = $1 where id = $2;

-- name: CreateScoreInstrument :exec
insert into score_instrument (score_id, instrument_id) values ($1, $2);

-- name: DeleteScoreInstrument :exec
delete from score_instrument where score_id = $1;

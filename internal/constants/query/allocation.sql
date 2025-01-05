-- name: GetAllocations :many
select * from allocation;

-- name: DeleteAllocation :exec
delete from allocation where id = $1;

-- name: CreateAllocation :one
insert into allocation (name) values ($1) returning *;

-- name: UpdateAllocation :exec
update allocation set name = $1 where id = $2;

-- name: CreateScoreAllocation :exec
insert into score_allocation (score_id, allocation_id) values ($1, $2);

-- name: DeleteScoreAllocation :exec
delete from score_allocation where score_id = $1;

-- name: GetCategories :many
select * from category;

-- name: DeleteCategory :exec
delete from category where id = $1;

-- name: CreateCategory :one
insert into category (name) values ($1) returning *;

-- name: UpdateCategory :exec
update category set name = $1 where id = $2;

-- name: CreateScoreCategory :exec
insert into score_category (score_id, category_id) values ($1, $2);

-- name: DeleteScoreCategory :exec
delete from score_category where score_id = $1; 

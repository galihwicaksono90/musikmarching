-- name: GetProfileById :one
select *
from profile p
where p.id = @id
limit 1
;

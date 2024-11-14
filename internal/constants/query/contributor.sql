-- name: GetContributorById :one
select 
  a.id,
  a.email,
  a.name,
  c.is_verified, 
  c.verified_at,
  c.created_at
from contributor as c
inner join account as a on c.id = a.id
where c.id = @id
;

-- name: CreateContributor :one
insert into contributor as c (id)
values (@id)
on conflict do nothing
returning c.id
;

-- name: GetUnverifiedContributors :many
select * from contributor as c
where c.is_verified = false;
;

-- name: VerifyContributor :exec
update contributor
set is_verified = true,
    verified_at = now()
where id = @id;
;

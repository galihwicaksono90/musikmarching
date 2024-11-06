-- name: GetContributorById :one
select 
  a.id,
  a.email,
  p.id, 
  p.name, 
  c.isverified, 
  c.verified_at,
  c.created_at
from contributor as c
inner join profile as p on c.id = p.id
inner join account as a on c.id = a.id
where p.id = @id
;

-- name: CreateContributor :one
insert into contributor as c (id)
values (@id)
on conflict do nothing
returning c.id
;

-- name: GetUnverifiedContributors :many
select * from contributor as c
where c.isverified = false;
;

-- name: VerifyContributor :exec
update contributor
set isverified = true,
    verifiedat = now()
where id = @id;
;

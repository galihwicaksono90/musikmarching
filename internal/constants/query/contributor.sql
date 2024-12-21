-- name: GetContributorById :one
select * from contributor_account_scores as cas
where cas.id = @id
;

-- name: CreateContributor :one
insert into contributor as c (id, full_name)
values (@id, @full_name)
on conflict do nothing
returning c.id
;

-- name: GetUnverifiedContributors :many
select * from contributor as c
where c.is_verified = false;
;

-- name: GetAllContributors :many
select * from contributor_account_scores as cas
;

-- name: VerifyContributor :exec
update contributor
set is_verified = true,
    verified_at = now()
where id = @id;
;

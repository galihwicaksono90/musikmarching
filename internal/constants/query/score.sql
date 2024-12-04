-- name: GetScores :many
select 
  s.id,
  s.title,
  s.is_verified,
  s.price,
  a.name,
  a.email
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a  on a.id = s.contributor_id
order by s.created_at desc
limit @pageLimit::int offset @pageOffset::int
; 


-- name: GetVerifiedScores :many
select 
  s.id,
  s.title,
  s.price,
  a.name,
  a.email
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a  on a.id = s.contributor_id
where s.is_verified = true and c.is_verified = true 
limit @pageLimit::int offset @pageOffset::int
; 


-- name: GetVerifiedScoreById :one
select 
  s.id,
  s.title,
  s.price,
  a.name as contributor_name
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a  on a.id = s.contributor_id
where s.is_verified = true and s.id = @id
; 

-- name: GetScoreById :one
select * from score s
where s.id = @id
; 

-- name: CreateScore :one
insert into score (
  title,
  price,
  pdf_url,
  music_url,
  contributor_id
) values (
  @title,
  @price,
  @pdf_url,
  @music_url,
  @contributor_id
) returning id;

-- name: UpdateScore :exec
update score set
  title = COALESCE(sqlc.narg('title'), title),
  price = COALESCE(sqlc.narg('price'), price),
  updated_at = now()
where id = @id
;

-- name: VerifyScore :exec
update score set
  is_verified = true,
  verified_at = now()
where id = @id
;

-- name: GetScoreByContributorId :many
select * from score
where contributor_id = @id
;



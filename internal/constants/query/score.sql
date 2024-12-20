-- name: GetScores :many
select s.id, s.title, s.is_verified, s.price, a.name, a.email
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
order by s.created_at desc
limit @pagelimit::int
offset @pageoffset::int
;

-- name: GetScoresPaginated :many
select *
from score s
where deleted_at is null
order by
    case when $3 = 'price_asc' then price when $3 = 'price_desc' then price end,
    case
        when $3 = 'created_at_asc'
        then created_at
        when $3 = 'created_at_desc'
        then created_at
    end desc
limit $1
offset $2
;


-- name: GetScoresByContributorID :many
select s.id, s.title, s.is_verified, s.price, a.name, a.email
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
where s.contributor_id = @id
order by s.is_verified desc, s.created_at desc
limit @pagelimit::int
offset @pageoffset::int
;

-- name: GetScoreByContributorID :one
select s.id, s.title, s.is_verified, s.price, a.name, a.email, s.pdf_url, s.audio_url, s.pdf_image_urls
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
where s.id = @score_id
and s.contributor_id = @contributor_id
;


-- name: GetVerifiedScores :many
select s.id, s.title, s.price, a.name, a.email
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
where s.is_verified = true and c.is_verified = true
limit @page_limit::int
offset @page_offset::int
;


-- name: GetVerifiedScoreById :one
select s.id, s.title, s.price, a.name as contributor_name
from score s
inner join contributor c on c.id = s.contributor_id
inner join account a on a.id = s.contributor_id
where s.is_verified = true and s.id = @id
;

-- name: GetScoreById :one
select *
from score s
where s.id = @id
;

-- name: CreateScore :one
insert into score (
  title,
  price,
  pdf_url,
  pdf_image_urls,
  audio_url,
  contributor_id
) values (
  @title,
  @price,
  @pdf_url,
  @pdf_image_urls,
  @audio_url,
  @contributor_id
) returning id;

-- name: UpdateScore :exec
update score set
  title = COALESCE(sqlc.narg('title'), title),
  price = COALESCE(sqlc.narg('price'), price),
  pdf_url = COALESCE(sqlc.narg('pdf_url'), pdf_url),
  pdf_image_urls = COALESCE(sqlc.narg('pdf_image_urls'), pdf_image_urls),
  audio_url = COALESCE(sqlc.narg('audio_url'), audio_url),
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
select *
from score
where contributor_id = @id
;


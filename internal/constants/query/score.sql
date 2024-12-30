-- name: GetAllPublicScores :many
select * from score_public_view spv
where spv.is_verified = true and spv.deleted_at is null
order by spv.created_at desc
limit @pagelimit::int
offset @pageoffset::int
;

-- select
--   s.id,
--   s.title,
--   s.is_verified,
--   s.price,
--   s.pdf_image_urls,
--   s.audio_url,
--   s.created_at,
--   a.email,
--   c.full_name,
--   COALESCE(ARRAY(SELECT i.name FROM instrument i
--                    JOIN score_instrument si ON i.id = si.instrument_id
--                    WHERE si.score_id = s.id
--                    ORDER BY i.name), ARRAY[]) AS instruments,
--   COALESCE(ARRAY(SELECT a.name FROM allocation a
--                    JOIN score_allocation sa ON a.id = sa.allocation_id
--                    WHERE sa.score_id = s.id
--                    ORDER BY a.name), ARRAY[]::TEXT[]) AS allocations,
--   COALESCE(ARRAY(SELECT c.name FROM category c
--                    JOIN score_category sc ON c.id = sc.category_id
--                    WHERE sc.score_id = s.id
--                    ORDER BY c.name), ARRAY[]::TEXT[]) AS categories
-- from score s
-- join contributor c on c.id = s.contributor_id
-- join account a on a.id = s.contributor_id
-- where s.deleted_at is null
-- and s.is_verified = true
-- and c.is_verified = true
-- order by s.created_at desc
-- limit @pagelimit::int
-- offset @pageoffset::int
-- ;


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
from score
where deleted_at is null
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


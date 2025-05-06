-- name: GetAllPublicScores :many
select sqlc.embed(spv), count(*) over() as count 
from score_public_view spv
where spv.is_verified = true and spv.deleted_at is null and spv.purchased_by is null
and (@title::text IS NULL OR lower(spv.title) like lower(@title))
and (sqlc.narg('difficulty')::difficulty IS NULL OR spv.difficulty in (sqlc.narg('difficulty')))
and (sqlc.narg('content_type')::content_type IS NULL OR spv.content_type in (sqlc.narg('content_type')))
and (@instruments::text[] IS NULL or spv.instruments::text[] && @instruments::text[])
and (@categories::text[] IS NULL or spv.categories::text[] && @categories::text[])
and (@allocations::text[] IS NULL or spv.allocations::text[] && @allocations::text[])
order by spv.created_at desc
limit @page_limit::int
offset @page_offset::int
;

-- name: GetPublicScoreById :one
select * from score_public_view spv
where spv.id = @id
and spv.is_verified = true 
and spv.deleted_at is null
and spv.purchased_by is null
limit 1
;

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
select scv.*, coalesce(p.purchase_count, 0)
from score_contributor_view scv
left join (
    select p.score_id, count(*) as purchase_count
    from purchase p
    inner join payment pm on pm.purchase_id = p.id
    group by p.score_id
) p on p.score_id = scv.id
where scv.contributor_id = @contributor_id::uuid
;

-- select * from score_contributor_view scv
-- where scv.contributor_id = @contributor_id
-- limit @pagelimit::int
-- offset @pageoffset::int
-- ;

-- name: GetScoreByContributorID :one
select * from score_contributor_view scv
where scv.id = @score_id
and scv.contributor_id = @contributor_id
limit 1
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
  contributor_id,
  description,
  content_type,
  difficulty
) values (
  @title,
  @price,
  @pdf_url,
  @pdf_image_urls,
  @audio_url,
  @contributor_id,
  @description,
  @content_type,
  @difficulty
) returning id;

-- name: UpdateScore :exec
update score set
  title = COALESCE(sqlc.narg('title'), title),
  price = COALESCE(sqlc.narg('price'), price),
  description = COALESCE(sqlc.narg('description'), description),
  difficulty = COALESCE(sqlc.narg('difficulty'), difficulty),
  content_type = COALESCE(sqlc.narg('content_type'), content_type),
  pdf_url = COALESCE(sqlc.narg('pdf_url'), pdf_url),
  pdf_image_urls = COALESCE(sqlc.narg('pdf_image_urls'), pdf_image_urls),
  audio_url = COALESCE(sqlc.narg('audio_url'), audio_url),
  updated_at = now()
where id = @id
;

-- name: DeleteScore :exec
update score s set
  deleted_at = now(),
  updated_at = now()
where s.id = @id 
and s.contributor_id = @contributor_id::uuid
;

-- name: VerifyScore :exec
update score set
  is_verified = true,
  verified_at = now()
where id = @id
;

-- name: GetScoreLibrary :many
with purchases as (
  select score_id from purchase  
  where account_id = @account_id::uuid
  and is_verified = true
)
select sqlc.embed(score_library_view),  count(*) over() as count
from score_library_view
where id in (select score_id from purchases)
limit @page_limit::int
offset @page_offset::int
;

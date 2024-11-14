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

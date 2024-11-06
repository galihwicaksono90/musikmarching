-- name: GetScoresByContributorId :many
select s.*
from contributor p
inner join contributor_score_uploads psu on p.id = psu.profile_id
inner join score s on psu.score_id = s.id
where p.id = @id
group by s.id, s.title;
;

-- name: CreateScore :exec
WITH score_insert AS (
  INSERT INTO Score (title)
  VALUES (@title)
  RETURNING id
)
INSERT INTO contributor_score_uploads (contributor_id, score_id)
values(@id, (SELECT id FROM score_insert))
;

-- name: GetVerifiedScores :many
select s.id, s.title, c.id, a.email, a.name from score as s
inner join contributor_score_uploads as csu on s.id = csu.score_id
inner join contributor as c on csu.contributor_id = c.id
inner join profile as p on p.id = c.id
inner join account as a on p.id = a.id
where c.isverified = true 
and s.isverified = true
;

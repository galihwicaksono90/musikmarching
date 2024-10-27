-- name: GetScoresByProfile :many
select s.*
from profile p
inner join profile_score_uploads psu on p.id = psu.profile_id
inner join score s on psu.score_id = s.id
where p.account_id = @account_id
group by s.id, s.title;
;

-- name: CreateScore :exec
WITH score_insert AS (
  INSERT INTO Score (title)
  VALUES (@title)
  RETURNING id
)
INSERT INTO profile_score_uploads (profile_id, score_id)
values(@id, (SELECT id FROM score_insert))
;

-- name: CreateContributorApply :one
insert into contributor_apply
(
  id,
  full_name,
  phone_number,
  musical_background,
  education,
  experience,
  portofolio_link,
  sample_url,
  terms_and_conditions_accepted,
  is_verified
)
values (
  @id,
  @full_name,
  @phone_number,
  @musical_background,
  @education,
  @experience,
  @portofolio_link,
  @sample_url,
  @terms_and_conditions_accepted,
  false
)
returning *
;

-- name: GetContributorApplyByAccountID :one
select * from contributor_apply where id = @account_id
;

-- name: UpdateContributorApply :exec
update contributor_apply set 
full_name = @full_name,
phone_number = @phone_number,
musical_background = COALESCE(sqlc.narg('musical_background'), musical_background),
education = COALESCE(sqlc.narg('education'), education),
experience = COALESCE(sqlc.narg('experience'), experience),
experience = COALESCE(sqlc.narg('portofolio_link'), portofolio_link),
updated_at = now()
where id = @account_id::uuid
;

-- name: GetContributorApplications :many
select * from contributor_apply
;

-- name: VerifyContributorApply :exec
update contributor_apply set 
is_verified = true,
updated_at = now()
where id = @account_id::uuid
;

-- name: CreateContributorFromContributorApply :exec
with contributor_apply_data as (
  select * from contributor_apply
  where id = @account_id::uuid
) 
insert into contributor(id, full_name, phone_number, musical_background, education, experience, portofolio_link, is_verified)
select id, full_name, phone_number, musical_background, education, experience, portofolio_link, true
from contributor_apply_data
;

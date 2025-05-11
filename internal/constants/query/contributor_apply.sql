-- name: CreateContributorApply :one
insert into contributor_apply
(
  account_id, 
  full_name,
  phone_number,
  musical_background,
  education,
  experience,
  portofolio_link,
  sample_url,
  terms_and_conditions_accepted
)
values (
  @account_id, 
  @full_name,
  @phone_number,
  @musical_background,
  @education,
  @experience,
  @portofolio_link,
  @sample_url,
  @terms_and_conditions_accepted
)
returning *
;

-- name: GetContributorApplyByAccountID :one
select * from contributor_apply where account_id = @account_id
;

-- name: UpdateContributorApply :exec
update contributor_apply set 
full_name = @full_name,
phone_number = @phone_number,
musical_background = COALESCE(sqlc.narg('musical_background'), musical_background),
education = COALESCE(sqlc.narg('education'), education),
experience = COALESCE(sqlc.narg('experience'), experience),
updated_at = now()
where account_id = @account_id::uuid
;

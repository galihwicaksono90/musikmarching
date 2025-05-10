-- name: GetAccountByEmail :one
select a.*, r.name as role_name from account a
inner join role r on r.id = a.role_id
where a.email = @email
limit 1
;

-- name: GetAccountById :one
select a.*, r.name as role_name from account a
inner join role r on r.id = a.role_id
where a.id = @id
limit 1
;

-- name: GetAccounts :many
select a.*, r.name as role_name from account a
inner join role r on r.id = a.role_id
;

-- name: CreateAccount :one
INSERT INTO Account (email, name, picture_url, role_id)
VALUES (@email, @name, @pictureurl, (select id from role where name = 'user'))
RETURNING id, name
;

-- name: UpdateAccount :one
update account a
set name = @name, picture_url = @pictureUrl
where id = @id
returning a.id
;

-- name: UpdateAccountRole :one
update account a
set role_id = (select id from role as r where r.name = @roleName)
where a.id = @id
RETURNING id
;

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
  sample_url
)
values (
  @account_id, 
  @full_name,
  @phone_number,
  @musical_background,
  @education,
  @experience,
  @portofolio_link,
  @sample_url
)
returning *
;


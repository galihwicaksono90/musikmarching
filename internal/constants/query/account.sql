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
WITH account_insert AS (
  INSERT INTO Account (email, name, picture_url, role_id)
  VALUES (@email, @name, @pictureurl, (select id from role where name = 'user'))
  RETURNING id
)
INSERT INTO Profile as p (account_id)
SELECT id FROM account_insert
returning account_id
;

-- name: UpdateAccount :one
update account a
set name = @name, picture_url = @pictureUrl
where id = @id
returning a.id
;

-- name: GetProfileByAccountId :one
select *
from profile p
where p.account_id = @account_id
limit 1
;

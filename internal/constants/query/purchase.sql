-- name: GetPurchaseByAccountAndScoreId :one
select * from purchase
where account_id = @account_id and score_id = @score_id
;

-- name: GetPurchaseById :one
select * from purchase
where id = @id and account_id = @account_id
;

-- name: CreatePurchase :one
insert into purchase (account_id, score_id, price, title)
values (@account_id, @score_id, @price, @title)
returning id
;

-- name: GetPurchasesByAccountId :many
select * from purchase
where account_id = @account_id
order by created_at desc
;

-- name: UpdatePurchaseProof :exec
update purchase set 
payment_proof_url = @payment_proof_url, 
paid_at = now(), updated_at = now()
where id = @id and account_id = @account_id
;

-- name: GetAllPurchases :many
select * from purchase
;

-- name: VerifyPurchase :exec
update purchase set
is_verified = true,
verified_at = now()
where id = @id
;


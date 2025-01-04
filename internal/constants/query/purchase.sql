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

-- name: GetPurchasedScoreById :one
select
    s.id,
    s.title,
    s.description,
    a.email,
    c.full_name,
    s.difficulty,
    s.content_type,
    s.pdf_url,
    s.pdf_image_urls,
    s.price,
    s.audio_url,
    s.is_verified
from score s
join contributor c on c.id = s.contributor_id
join account a on a.id = c.id
where s.id = @score_id and s.is_verified = true
limit 1
;

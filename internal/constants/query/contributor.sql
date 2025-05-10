-- name: GetContributorById :one
select * from contributor_account_scores as cas
where cas.id = @id
;

-- name: CreateContributor :one
insert into contributor as c (id, full_name)
values (@id, @full_name)
on conflict do nothing
returning c.id
;

-- name: GetUnverifiedContributors :many
select * from contributor as c
where c.is_verified = false;
;

-- name: GetAllContributors :many
select * from contributor_account_scores as cas
;

-- name: VerifyContributor :exec
update contributor
set is_verified = true,
    verified_at = now()
where id = @id;
;

-- name: GetContributorScoreStatistics :one
select 
  sum(py.revenue) as revenue, 
  count(p.*) as purchase_count, 
  (
  select count(*) from score 
    where contributor_id = @contributor_id::uuid
  ) as score_count
from purchase p
  inner join score s on p.score_id = s.id 
  inner join payment py on py.purchase_id = p.id
where p.is_verified = true 
  and s.contributor_id = @contributor_id::uuid
;
 
-- name: GetContributorBestSellingScores :many
select s.id, s.title, count(p.score_id) as count, sum(py.revenue) as revenue 
from purchase p
inner join score s on p.score_id = s.id
inner join payment py on py.purchase_id = p.id
where s.contributor_id = @contributor_id::uuid
group by p.score_id, s.id
order by count desc
limit 5
;

-- name: GetContributorPaymentMethod :one
select p.* from payment_method p
inner join contributor c on p.contributor_id = c.id
where c.id = @contributor_id::uuid
limit 1
;

-- name: UpsertContributorPaymentMethod :exec
insert into payment_method (method, account_number, contributor_id, bank_name)
values (@method, @account_number, @contributor_id::uuid, @bank_name)
  on conflict (contributor_id)
do update set method = @method,
              account_number = @account_number,
              bank_name = @bank_name,
              updated_at = now()
;

-- name: GetContributorPayments :many
select 
  py.id, 
  py.purchase_id, 
  py.price, 
  py.revenue, 
  py.created_at, 
  py.is_verified,
  py.payment_method,
  py.account_number,
  py.bank_name,
  pu.invoice_serial, 
  c.id as contributor_id,
  s.title  
from payment py
inner join purchase pu on py.purchase_id = pu.id
inner join score s on s.id = pu.score_id
inner join contributor c on c.id = s.contributor_id
where c.id = @contributor_id::uuid
;

-- name: GetContributorPaymentStatistics :one
select 
    sum(py.revenue) as total_paid,
    (
      select py.verified_at from payment py
      inner join purchase pu on py.purchase_id = pu.id
      inner join score s on s.id = pu.score_id
      where pu.is_verified = true
      and s.contributor_id = @contributor_id::uuid
      order by py.verified_at desc
      limit 1
    ) as latest_payment
from payment py
inner join purchase pu on py.purchase_id = pu.id
inner join score s on pu.score_id = s.id
and s.contributor_id = @contributor_id::uuid
;

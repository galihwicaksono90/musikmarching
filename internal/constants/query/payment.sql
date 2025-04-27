-- name: CreatePayment :exec
with pur as (
  select id, price, score_id from purchase
  where id = @purchase_id::uuid
),
scr as (
  select contributor_id from score
  where id = (select score_id from pur)
),
cpm as (
  select method, bank_name, account_number from payment_method
  where contributor_id = (select contributor_id from scr)
)
insert into payment
(purchase_id, price, revenue, fee_percentage, payment_method, account_number, bank_name)
select 
  pur.id, 
  pur.price, 
  pur.price - (pur.price * (20.00 / 100.00)),
  20.00, 
  cpm.method, 
  cpm.account_number, 
  cpm.bank_name
from pur cross join cpm;

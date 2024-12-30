-- +goose Up
-- +goose StatementBegin
insert into role as r (id, name)
values 
  ('515ebadd-f8d1-472f-8b8c-1ba53d61a358', 'admin'),
  ('78653d16-6134-4f84-afb6-f44deb51f898', 'contributor'),
  ('748d4130-fc58-4485-be80-f49342252132', 'user')
;

insert into account (id, email, name, picture_url, role_id)
values
(
  '291a7f36-69ab-4be1-ba91-064445349bbd', 
  'gorillahobo@gmail.com', 
  'gorillahobo', 
  'https://lh3.googleusercontent.com/a/ACg8ocJ5FwcGkTjLYTdbgvTsbKdQzVQaHcytNSMJHKFkBPptwbjRu-c=s96-c',
  (select id from role where name = 'admin')
),
(
  'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 
  'galihwicaksono90@gmail.com', 
  'Galih Wicaksono', 
  'https://lh3.googleusercontent.com/a/ACg8ocJ5FwcGkTjLYTdbgvTsbKdQzVQaHcytNSMJHKFkBPptwbjRu-c=s96-c',
  (select id from role where name = 'user')
),
(
  'ab48aeb7-51a1-4712-932b-fe64d98fec87', 
  'galih.wicaksono@softwareseni.com', 
  'Galih Wicaksono', 
  'https://lh3.googleusercontent.com/a/ACg8ocJWwtbtedC3Ys49-9UFEcTJ4xiUFAdTLikFNOwuemfqxPaIEYE=s96-c',
  (select id from role where name = 'contributor')
)
;

insert into contributor (id, full_name, is_verified, verified_at)
values ('ab48aeb7-51a1-4712-932b-fe64d98fec87', 'Galih Wicaksono', true, now());

insert into score (id, title, price, difficulty, contributor_id, pdf_url, pdf_image_urls, audio_url, is_verified, verified_at)
values 
  ('43a89d84-706c-4727-b1d8-191731bce558', 'song one',   100.00, 'beginner', 'ab48aeb7-51a1-4712-932b-fe64d98fec87', 'pdf_url', '{hello, world}', 'audio_url', true,  now()),
  ('97320ce1-3159-4ccd-a645-ba7e80f03a5a', 'song two',   200.00, 'intermediate', 'ab48aeb7-51a1-4712-932b-fe64d98fec87', 'pdf_url', '{hello, world}', 'audio_url', false, null ),
  ('a31d7a7d-4e85-4b27-b956-2bff6415ddd6', 'song three', 300.00, 'advanced', 'ab48aeb7-51a1-4712-932b-fe64d98fec87', 'pdf_url', '{hello, world}', 'audio_url', true,  now()),
  ('c17db1c2-2c35-477b-8dd7-00fb6db9723c', 'song four',  400.00, 'beginner', 'ab48aeb7-51a1-4712-932b-fe64d98fec87', 'pdf_url', '{hello, world}', 'audio_url', false, null )
;

insert into instrument (name)
values 
  ('Pianica'),
  ('Brass'),
  ('Other')
;

insert into score_instrument (score_id, instrument_id)
values
  ('43a89d84-706c-4727-b1d8-191731bce558', 1),
  ('43a89d84-706c-4727-b1d8-191731bce558', 2),
  ('a31d7a7d-4e85-4b27-b956-2bff6415ddd6', 1),
  ('c17db1c2-2c35-477b-8dd7-00fb6db9723c', 2)
;

insert into allocation (name)
values 
  ('SD'),
  ('SMP'),
  ('Umum')
;

insert into score_allocation (score_id, allocation_id)
values
  ('43a89d84-706c-4727-b1d8-191731bce558', 1),
  ('43a89d84-706c-4727-b1d8-191731bce558', 2),
  ('a31d7a7d-4e85-4b27-b956-2bff6415ddd6', 1),
  ('c17db1c2-2c35-477b-8dd7-00fb6db9723c', 2)
;

insert into category (name)
values 
  ('Ensemble/Concert'),
  ('Display'),
  ('Parade & Show')
;

insert into score_category (score_id, category_id)
values
  ('43a89d84-706c-4727-b1d8-191731bce558', 1),
  ('43a89d84-706c-4727-b1d8-191731bce558', 2),
  ('a31d7a7d-4e85-4b27-b956-2bff6415ddd6', 1),
  ('c17db1c2-2c35-477b-8dd7-00fb6db9723c', 2)
;



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd

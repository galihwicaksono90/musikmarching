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
  'b5014504-c91d-4f20-8ef7-22d63ee67443', 
  'adika.pperdana@gmail.com', 
  'Admin', 
  'https://lh3.googleusercontent.com/a/ACg8ocJ5FwcGkTjLYTdbgvTsbKdQzVQaHcytNSMJHKFkBPptwbjRu-c=s96-c',
  (select id from role where name = 'admin')
),
(
  'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 
  'galihwicaksono90@gmail.com', 
  'Galih Wicaksono', 
  'https://lh3.googleusercontent.com/a/ACg8ocJ5FwcGkTjLYTdbgvTsbKdQzVQaHcytNSMJHKFkBPptwbjRu-c=s96-c',
  (select id from role where name = 'contributor')
),
(
  'ab48aeb7-51a1-4712-932b-fe64d98fec87', 
  'galih.wicaksono@softwareseni.com', 
  'Galih Wicaksono', 
  'https://lh3.googleusercontent.com/a/ACg8ocJWwtbtedC3Ys49-9UFEcTJ4xiUFAdTLikFNOwuemfqxPaIEYE=s96-c',
  (select id from role where name = 'user')
),
(
  '52240b2c-ec89-4415-89de-ef4b35078486', 
  'swaranadamusic.email@gmail.com', 
  'Swaranada Music', 
  'https://lh3.googleusercontent.com/a/ACg8ocJWwtbtedC3Ys49-9UFEcTJ4xiUFAdTLikFNOwuemfqxPaIEYE=s96-c',
  (select id from role where name = 'contributor')
)
;

insert into contributor (id, full_name, is_verified, verified_at)
values 
  ('f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'Tony Blank', true, now()),
  ('52240b2c-ec89-4415-89de-ef4b35078486', 'Swaranada Musik', true, now())
;

insert into score (id, title, description, price, difficulty, contributor_id, pdf_url, pdf_image_urls, audio_url, is_verified, verified_at, content_type)
values 
  ('43a89d84-706c-4727-b1d8-191731bce558', 'Song One', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 100.00, 'beginner', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', true,  now(), 'exclusive'),
  ('97320ce1-3159-4ccd-a645-ba7e80f03a5a', 'Song Two', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 200.00, 'intermediate', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', false, null, 'non-exclusive' ),
  ('a31d7a7d-4e85-4b27-b956-2bff6415ddd6', 'Song Three', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 300.00, 'advanced', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', true,  now(), 'non-exclusive'),
  ('c17db1c2-2c35-477b-8dd7-00fb6db9723c', 'Song Four', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 400.00, 'beginner', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', false, null, 'exclusive'),
  ('2bff6910-66be-414c-ad73-f66413f12e41', 'Song Five', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 400.00, 'intermediate', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', true, now(), 'non-exclusive'),
  ('7c2359ff-8ef2-4750-abeb-db923e215450', 'Song Six', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 400.00, 'beginner', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', false, null, 'exclusive'),
  ('d70f4fe7-b2a0-49a3-8231-f678fd350ff8', 'Song Seven', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 400.00, 'intermediate', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', true, now(), 'non-exclusive'),
  ('0d37c6b4-3381-46b7-9763-8eb66f848876', 'Song Eight', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 400.00, 'beginner', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', false, null, 'exclusive'),
  ('8931667c-aa8d-4786-8b22-8a7546ba9433', 'Song Nine', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 400.00, 'advanced', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', true, now(), 'non-exclusive'),
  ('f1dfb05b-b9a0-4d15-94b1-ed75f9e2d59b', 'Song Ten', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc malesuada venenatis enim, vel mollis felis mollis tincidunt. Curabitur aliquam est ut mauris vulputate, eget feugiat lacus commodo. Integer quis nisl pretium, laoreet quam at, tristique dui. Maecenas et ligula vestibulum, pellentesque ante sit amet, semper eros. Sed a ligula magna. Ut convallis, velit eget congue laoreet, purus massa hendrerit ex, eu vehicula tellus elit eget mauris. Aenean molestie volutpat diam, vel lacinia velit molestie vitae', 400.00, 'beginner', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 'https://s2.q4cdn.com/175719177/files/doc_presentations/Placeholder-PDF.pdf', '{https://picsum.photos/200/300?random=1, https://picsum.photos/200/300?random=2}', 'https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3', false, null, 'exclusive')
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

insert into payment_method(id, method, bank_name, account_number, contributor_id)
values 
  ('beaf7202-09d9-4d71-a985-dfa17f8221a5', 'Transfer', 'BCA', '123456789', 'f45cb09c-7ef3-473a-a8df-0e580ad026d1'),
  ('d28ddde8-d8ee-4c2b-9399-f7b689b9945a', 'Transfer', 'BCA', '123456789', '52240b2c-ec89-4415-89de-ef4b35078486')
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE Score (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  contributor_id UUID NOT NULL references contributor (id),
  title VARCHAR(255) NOT NULL,
  price decimal(16,2) NOT NULL,
  is_verified BOOLEAN NOT NULL DEFAULT false,
  verified_at TIMESTAMPTZ,
  pdf_url VARCHAR(255),
  music_url VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

alter table score
add CONSTRAINT fk_score_contributor 
FOREIGN KEY(contributor_id) REFERENCES contributor(id) deferrable initially deferred;

insert into score (id, title, price, contributor_id, is_verified, verified_at)
values 
  ('43a89d84-706c-4727-b1d8-191731bce558', 'song one',   100.00, 'ab48aeb7-51a1-4712-932b-fe64d98fec87', true,  now()),
  ('97320ce1-3159-4ccd-a645-ba7e80f03a5a', 'song two',   200.00, 'ab48aeb7-51a1-4712-932b-fe64d98fec87', false, null ),
  ('a31d7a7d-4e85-4b27-b956-2bff6415ddd6', 'song three', 300.00, 'ab48aeb7-51a1-4712-932b-fe64d98fec87', true,  now()),
  ('c17db1c2-2c35-477b-8dd7-00fb6db9723c', 'song four',  400.00, 'ab48aeb7-51a1-4712-932b-fe64d98fec87', false, null )
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Score cascade;
-- +goose StatementEnd

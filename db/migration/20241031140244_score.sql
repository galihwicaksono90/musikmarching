-- +goose Up
-- +goose StatementBegin
CREATE TABLE Score (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  isVerified BOOLEAN NOT NULL DEFAULT false,
  verifiedAt TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

insert into score (id, title, isVerified, verifiedAt)
values ('43a89d84-706c-4727-b1d8-191731bce558', 'song one', true, now());
insert into score (id, title)
values ('97320ce1-3159-4ccd-a645-ba7e80f03a5a', 'song two');
insert into score (id, title, isVerified, verifiedAt)
values ('a31d7a7d-4e85-4b27-b956-2bff6415ddd6', 'song three', true, now());
insert into score (id, title)
values ('c17db1c2-2c35-477b-8dd7-00fb6db9723c', 'song four');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Score cascade;
-- +goose StatementEnd

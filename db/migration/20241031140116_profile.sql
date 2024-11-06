-- +goose Up
-- +goose StatementBegin
CREATE TABLE Profile (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255),
  CONSTRAINT fk_account FOREIGN KEY(id) REFERENCES account(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

insert into profile (id, name)
values('291a7f36-69ab-4be1-ba91-064445349bbd', (select name from account where id = '291a7f36-69ab-4be1-ba91-064445349bbd'));
insert into profile (id, name)
values('f45cb09c-7ef3-473a-a8df-0e580ad026d1', (select name from account where id = 'f45cb09c-7ef3-473a-a8df-0e580ad026d1'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Profile cascade;
-- +goose StatementEnd

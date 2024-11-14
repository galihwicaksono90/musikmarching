-- +goose Up
-- +goose StatementBegin
CREATE TABLE Account (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  picture_url VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,
  role_id UUID NOT NULL REFERENCES Role(id),

  UNIQUE(email)
);

alter table account
add constraint fk_account_role 
FOREIGN KEY (role_id) REFERENCES role (id) deferrable initially deferred;

insert into account (id, email, name, picture_url, role_id)
values
(
  '291a7f36-69ab-4be1-ba91-064445349bbd', 
  'gorillahobo@gmail.com', 
  'gorillahobo', 
  'https://lh3.googleusercontent.com/a/ACg8ocJ5FwcGkTjLYTdbgvTsbKdQzVQaHcytNSMJHKFkBPptwbjRu-c=s96-c',
  (select id from role where name = 'contributor')
),
(
  'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 
  'galihwicaksono90@gmail.com', 
  'Galih Wicaksono', 
  'https://lh3.googleusercontent.com/a/ACg8ocJ5FwcGkTjLYTdbgvTsbKdQzVQaHcytNSMJHKFkBPptwbjRu-c=s96-c',
  (select id from role where name = 'user')
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Account cascade;
-- +goose StatementEnd

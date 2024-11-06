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

insert into account (id, email, name, picture_url, role_id)
values (
  'f45cb09c-7ef3-473a-a8df-0e580ad026d1', 
  'galihwicaksono90@gmail.com', 
  'Galih Wicaksono', 
  'https://lh3.googleusercontent.com/a/ACg8ocJ5FwcGkTjLYTdbgvTsbKdQzVQaHcytNSMJHKFkBPptwbjRu-c=s96-c',
  '515ebadd-f8d1-472f-8b8c-1ba53d61a358'
)
;

insert into account (id, email, name, role_id)
values (
  '291a7f36-69ab-4be1-ba91-064445349bbd', 
  'gorillahobo@gmail.com', 
  'gorillahobo', 
  (select id from role where name = 'contributor')
)
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Account cascade;
-- +goose StatementEnd

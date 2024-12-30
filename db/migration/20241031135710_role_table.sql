-- +goose Up
-- +goose StatementBegin
CREATE TYPE RoleName AS ENUM ('admin', 'contributor', 'user');

CREATE TABLE Role (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name RoleName NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,

  UNIQUE(name)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Role cascade;
DROP type RoleName cascade;
-- +goose StatementEnd

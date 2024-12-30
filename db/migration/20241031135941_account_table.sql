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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Account cascade;
-- +goose StatementEnd

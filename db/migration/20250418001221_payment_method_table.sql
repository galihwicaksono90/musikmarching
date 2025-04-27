-- +goose Up
-- +goose StatementBegin
CREATE TABLE payment_method (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  method VARCHAR(255) NOT NULL,
  bank_name VARCHAR(255) NOT NULL,
  account_number VARCHAR(255) NOT NULL,
  contributor_id UUID NOT NULL references contributor (id) deferrable initially deferred unique,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payment_method;
-- +goose StatementEnd

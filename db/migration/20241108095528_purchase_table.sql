-- +goose Up
-- +goose StatementBegin
CREATE TABLE Purchase (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  invoice_serial serial not null,
  account_id UUID NOT NULL references account (id) deferrable initially deferred,
  score_id UUID NOT NULL references score (id) deferrable initially deferred,
  price DECIMAL(10,2) NOT NULL,
  title VARCHAR(255) NOT NULL,
  payment_proof_url VARCHAR(255),
  paid_at TIMESTAMPTZ,
  is_verified BOOLEAN NOT NULL DEFAULT false,
  verified_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Purchase cascade;
-- +goose StatementEnd

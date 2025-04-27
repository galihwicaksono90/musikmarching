-- +goose Up
-- +goose StatementBegin
CREATE TABLE Payment (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  purchase_id UUID NOT NULL references purchase (id) deferrable initially deferred unique,
  price DECIMAL(10,2) NOT NULL,
  revenue DECIMAL(10,2) NOT NULL,
  fee_percentage INTEGER NOT NULL,
  payment_proof_url VARCHAR(255),
  payment_method VARCHAR(255),
  account_number VARCHAR(255),
  bank_name VARCHAR(255),
  paid_at TIMESTAMPTZ,
  is_verified BOOlEAN NOT NULL DEFAULT false,
  verified_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,
  FOREIGN KEY(purchase_id) REFERENCES purchase(id) on delete cascade
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Payment cascade;
-- +goose StatementEnd

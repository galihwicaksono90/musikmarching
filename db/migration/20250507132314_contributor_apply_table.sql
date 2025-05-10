-- +goose Up
-- +goose StatementBegin
CREATE TABLE contributor_apply (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id UUID NOT NULL references account (id) deferrable initially deferred unique,
  full_name VARCHAR(255) NOT NULL,
  phone_number VARCHAR(255),
  musical_background TEXT NOT NULL,
  education TEXT,
  experience TEXT,
  portofolio_link TEXT,
  sample_url VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE contributor_apply;
-- +goose StatementEnd

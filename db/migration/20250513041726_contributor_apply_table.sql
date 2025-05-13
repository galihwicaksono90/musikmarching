-- +goose Up
-- +goose StatementBegin
CREATE TABLE contributor_apply (
  id UUID PRIMARY KEY references account (id),
  is_verified BOOLEAN NOT NULL DEFAULT false,
  full_name VARCHAR(255) NOT NULL,
  phone_number VARCHAR(255) NOT NULL,
  musical_background TEXT NOT NULL,
  education TEXT,
  experience TEXT,
  portofolio_link TEXT,
  terms_and_conditions_accepted BOOLEAN NOT NULL DEFAULT false,
  sample_url VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE contributor_apply;
-- +goose StatementEnd

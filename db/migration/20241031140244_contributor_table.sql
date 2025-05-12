-- +goose Up
-- +goose StatementBegin
CREATE TABLE Contributor (
  id UUID PRIMARY KEY references account (id),
  is_verified BOOLEAN DEFAULT false,
  full_name TEXT NOT NULL,
  verified_at TIMESTAMPTZ,
  phone_number VARCHAR(255) NOT NULL,
  musical_background TEXT NOT NULL,
  education TEXT,
  experience TEXT,
  portofolio_link TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

alter table contributor
add CONSTRAINT fk_contributor_account 
FOREIGN KEY (id) REFERENCES account (id) deferrable initially deferred;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Contributor cascade;
-- +goose StatementEnd

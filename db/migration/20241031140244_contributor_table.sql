-- +goose Up
-- +goose StatementBegin
CREATE TABLE Contributor (
  id UUID PRIMARY KEY references account (id),
  is_verified BOOLEAN DEFAULT false,
  verified_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

alter table contributor
add CONSTRAINT fk_contributor_account 
FOREIGN KEY (id) REFERENCES account (id) deferrable initially deferred;

insert into contributor (id, is_verified, verified_at)
values ('291a7f36-69ab-4be1-ba91-064445349bbd', true, now());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Contributor cascade;
-- +goose StatementEnd

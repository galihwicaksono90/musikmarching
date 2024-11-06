-- +goose Up
-- +goose StatementBegin
CREATE TABLE Contributor (
  id UUID PRIMARY KEY,
  isVerified BOOLEAN DEFAULT false,
  verified_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

insert into contributor (id)
values ('291a7f36-69ab-4be1-ba91-064445349bbd');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Contributor cascade;
-- +goose StatementEnd

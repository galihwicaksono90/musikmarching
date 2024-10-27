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

insert into role as r (id, name)
values 
  ('515ebadd-f8d1-472f-8b8c-1ba53d61a358', 'admin'),
  ('78653d16-6134-4f84-afb6-f44deb51f898', 'contributor'),
  ('748d4130-fc58-4485-be80-f49342252132', 'user')
;

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

CREATE TABLE Profile (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id UUID NOT NULL,
  CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES account(id)
);

CREATE TABLE Score (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

CREATE TABLE Profile_Score_Uploads (
  profile_id UUID,
  score_id UUID,
  PRIMARY KEY (profile_id, score_id),
  CONSTRAINT fk_profile FOREIGN KEY(profile_id) REFERENCES profile(id),
  CONSTRAINT fk_score FOREIGN KEY(score_id) REFERENCES score(id)
);


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE Role cascade;
DROP TABLE Account cascade;
DROP TABLE Profile cascade;
DROP type RoleName cascade;
DROP TABLE Score cascade;
DROP TABLE Profile_score_uploads cascade;
-- +goose StatementEnd



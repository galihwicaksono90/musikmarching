-- +goose Up
-- +goose StatementBegin
CREATE TYPE difficulty AS ENUM ('beginner', 'intermediate', 'advanced');
CREATE TYPE content_type AS ENUM ('exclusive', 'non-exclusive');

CREATE TABLE Score (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  contributor_id UUID NOT NULL references contributor (id),
  title VARCHAR(255) NOT NULL,
  description Text,
  price decimal(16,2) NOT NULL,
  is_verified BOOLEAN NOT NULL DEFAULT false,
  content_type content_type NOT NULL default 'non-exclusive',
  purchased_by UUID references account(id),
  verified_at TIMESTAMPTZ,
  difficulty Difficulty NOT NULL,
  pdf_url VARCHAR(255) NOT NULL,
  pdf_image_urls VARCHAR(255)[] NOT NULL DEFAULT '{}',
  audio_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,
  FOREIGN KEY(purchased_by) REFERENCES account(id) on delete cascade
);

alter table score
add CONSTRAINT fk_score_contributor 
FOREIGN KEY(contributor_id) REFERENCES contributor(id) deferrable initially deferred
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Score cascade;
DROP type difficulty cascade;
DROP type content_type cascade;
-- +goose StatementEnd

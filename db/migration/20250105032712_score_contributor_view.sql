-- +goose Up
-- +goose StatementBegin
CREATE or replace VIEW score_contributor_view AS
select
  s.id,
  s.title,
  s.description,
  s.is_verified,
  s.price,
  s.difficulty,
  s.content_type,
  s.purchased_by,
  s.pdf_image_urls,
  s.pdf_url,
  s.audio_url,
  s.created_at,
  s.updated_at,
  s.deleted_at,
  a.email,
  c.full_name,
  c.id as contributor_id,
  COALESCE(ARRAY(SELECT i.id FROM instrument i
                   JOIN score_instrument si ON i.id = si.instrument_id
                   WHERE si.score_id = s.id
                   ORDER BY i.name), ARRAY[]::INT[]) AS instruments,
  COALESCE(ARRAY(SELECT a.id FROM allocation a
                   JOIN score_allocation sa ON a.id = sa.allocation_id
                   WHERE sa.score_id = s.id
                   ORDER BY a.name), ARRAY[]::INT[]) AS allocations,
  COALESCE(ARRAY(SELECT c.id FROM category c
                   JOIN score_category sc ON c.id = sc.category_id
                   WHERE sc.score_id = s.id
                   ORDER BY c.name), ARRAY[]::INT[]) AS categories
from score s
join contributor c on c.id = s.contributor_id
join account a on a.id = s.contributor_id
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view score_contributor_view;
-- +goose StatementEnd

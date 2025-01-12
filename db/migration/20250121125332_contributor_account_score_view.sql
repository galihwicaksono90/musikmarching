-- +goose Up
-- +goose StatementBegin
CREATE or replace VIEW contributor_account_scores AS
SELECT
    c.*,
    a.email,
    COALESCE(JSON_AGG(s.*) FILTER (WHERE s.contributor_id IS NOT NULL), '[]') as scores
FROM
    contributor c
JOIN 
    account a on a.id = c.id
LEFT JOIN
    score_contributor_view s ON s.contributor_id = c.id
GROUP BY
    c.id, c.full_name, a.email;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view contributor_account_scores;
-- +goose StatementEnd

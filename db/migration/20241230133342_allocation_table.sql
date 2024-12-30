-- +goose Up
-- +goose StatementBegin
CREATE TABLE allocation (
  id SERIAL PRIMARY KEY, 
  name VARCHAR(255) NOT NULL UNIQUE
)
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE allocation cascade;
-- +goose StatementEnd

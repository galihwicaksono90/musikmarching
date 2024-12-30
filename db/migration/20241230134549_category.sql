-- +goose Up
-- +goose StatementBegin
CREATE TABLE category (
  id SERIAL PRIMARY KEY, 
  name VARCHAR(255) NOT NULL UNIQUE
)
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE category cascade;
-- +goose StatementEnd

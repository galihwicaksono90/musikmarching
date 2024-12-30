-- +goose Up
-- +goose StatementBegin
CREATE TABLE score_category (
    score_id UUID NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY (score_id, category_id),
    FOREIGN KEY (score_id) REFERENCES score(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
)
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE score_category cascade;
-- +goose StatementEnd

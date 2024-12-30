-- +goose Up
-- +goose StatementBegin
CREATE TABLE score_instrument (
    score_id UUID NOT NULL,
    instrument_id INT NOT NULL,
    PRIMARY KEY (score_id, instrument_id),
    FOREIGN KEY (score_id) REFERENCES score(id) ON DELETE CASCADE,
    FOREIGN KEY (instrument_id) REFERENCES instrument(id) ON DELETE CASCADE
)
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE score_instrument cascade;
-- +goose StatementEnd

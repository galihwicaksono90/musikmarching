-- +goose Up
-- +goose StatementBegin
CREATE TABLE score_allocation (
    score_id UUID NOT NULL,
    allocation_id INT NOT NULL,
    PRIMARY KEY (score_id, allocation_id),
    FOREIGN KEY (score_id) REFERENCES score(id) ON DELETE CASCADE,
    FOREIGN KEY (allocation_id) REFERENCES allocation(id) ON DELETE CASCADE
)
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE score_allocation cascade;
-- +goose StatementEnd

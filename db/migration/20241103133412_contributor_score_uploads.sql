-- +goose Up
-- +goose StatementBegin
CREATE TABLE contributor_Score_Uploads (
  contributor_id UUID,
  score_id UUID,
  PRIMARY KEY (contributor_id, score_id),
  CONSTRAINT fk_contributor FOREIGN KEY(contributor_id) REFERENCES profile(id),
  CONSTRAINT fk_score FOREIGN KEY(score_id) REFERENCES score(id)
);

insert into contributor_score_uploads (contributor_id, score_id)
values ('291a7f36-69ab-4be1-ba91-064445349bbd', '43a89d84-706c-4727-b1d8-191731bce558');
insert into contributor_score_uploads (contributor_id, score_id)
values ('291a7f36-69ab-4be1-ba91-064445349bbd', '97320ce1-3159-4ccd-a645-ba7e80f03a5a');
insert into contributor_score_uploads (contributor_id, score_id)
values ('291a7f36-69ab-4be1-ba91-064445349bbd', 'a31d7a7d-4e85-4b27-b956-2bff6415ddd6');
insert into contributor_score_uploads (contributor_id, score_id)
values ('291a7f36-69ab-4be1-ba91-064445349bbd', 'c17db1c2-2c35-477b-8dd7-00fb6db9723c');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE contributor_score_uploads cascade;
-- +goose StatementEnd

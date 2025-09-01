-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE question_option(
  id BIGSERIAL PRIMARY KEY,
  question_id BIGINT NOT NULL REFERENCES question(id) ON DELETE CASCADE,
  text text NOT NULL,
  is_correct boolean NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE question_option
-- +goose StatementEnd

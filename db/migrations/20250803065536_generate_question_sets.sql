-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO question_set(question_id, set_id)
SELECT
    question_id,
    (random() * 9)::int + 1 AS set_id
FROM generate_series(1, 100) AS question_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

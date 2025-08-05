-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO question_set(name, created_at)
SELECT
    md5(random()::text) AS name,
    now() - (random() * INTERVAL '365 days') AS created_at
FROM generate_series(1, 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

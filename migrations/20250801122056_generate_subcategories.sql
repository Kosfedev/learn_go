-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO subcategory(name, category_id, created_at)
SELECT
    md5(random()::text) as name,
    (random() * 99) + 1 as category_id,
    now() - (random() * INTERVAL '365 days') AS created_at
FROM generate_series(1, 1000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

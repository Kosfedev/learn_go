-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO category(name, domain_id, created_at)
SELECT
    md5(random()::text) AS name,
    (random() * 9)::int + 1 AS domain_id,
    now() - ((random()+1) * INTERVAL '365 days') AS created_at
FROM generate_series(1, 100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

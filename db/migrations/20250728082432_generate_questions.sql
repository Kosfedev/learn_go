-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
WITH
    questions(text) AS (
    VALUES
        ('Explain quantum entanglement in simple terms'),
        ('How do neural networks learn?'),
        ('What is the meaning of life?'),
        ('Describe photosynthesis process'),
        ('Python vs Go: which is better for web?')
),
    data AS (
    SELECT
        (random() * 2 + 1)::int AS type,
        now() - (random() * INTERVAL '365 days') AS created_at
    FROM generate_series(1, 200)
)
INSERT INTO question (text, type, reference_answer, created_at, updated_at)
SELECT
    questions.text AS text,
    data.type AS type,
    CASE WHEN data.type = 3 OR random() >= 0.7
        THEN md5(random()::text)
        END AS reference_answer,
    data.created_at AS created_at,
    CASE WHEN random() < 0.3
        THEN data.created_at + (random() * (now() - data.created_at))
        END AS updated_at
FROM data
CROSS JOIN questions;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

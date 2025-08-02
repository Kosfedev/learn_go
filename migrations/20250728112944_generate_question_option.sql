-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO question_option(question_id, text, is_correct)
SELECT
    (random() * 9)::int + 1 as question_id,
    md5(random()::TEXT) as text,
    CASE WHEN random()::int = 1
        THEN True
        ELSE FALSE
        END as is_correct
FROM generate_series(1, 50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

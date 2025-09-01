-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO question_subcategory(question_id, subcategory_id)
SELECT
    question_id,
    (random() * 9)::int + 1 AS subcategory_id
FROM generate_series(1,100) AS question_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

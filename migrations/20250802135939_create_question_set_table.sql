-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE question_set(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE question_set;
-- +goose StatementEnd

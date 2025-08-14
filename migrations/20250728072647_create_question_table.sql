-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE question (
    id BIGSERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    type INTEGER,
    reference_answer TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE question;
-- +goose StatementEnd

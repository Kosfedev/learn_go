-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE category(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    domain_id BIGINT NOT NULL REFERENCES domain(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE category;
-- +goose StatementEnd

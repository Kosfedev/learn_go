-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE subcategory(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    category_id BIGINT NOT NULL REFERENCES category(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE subcategory;
-- +goose StatementEnd

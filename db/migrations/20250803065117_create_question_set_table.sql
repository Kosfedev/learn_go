-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE question_set(
    question_id BIGINT NOT NULL,
    set_id BIGINT NOT NULL,

    PRIMARY KEY (question_id, set_id),
    FOREIGN KEY (question_id) REFERENCES question(id) ON DELETE CASCADE,
    FOREIGN KEY (set_id) REFERENCES set(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE question_set;
-- +goose StatementEnd

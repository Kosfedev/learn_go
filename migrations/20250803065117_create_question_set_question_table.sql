-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE question_set_question(
    question_id INTEGER NOT NULL,
    question_set_id INTEGER NOT NULL,

    PRIMARY KEY (question_id, question_set_id),
    FOREIGN KEY (question_id) REFERENCES question(id) ON DELETE CASCADE,
    FOREIGN KEY (question_set_id) REFERENCES question_set(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE question_set_question;
-- +goose StatementEnd

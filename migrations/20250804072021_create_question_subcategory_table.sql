-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE question_subcategory(
    question_id INTEGER NOT NULL,
    subcategory_id INTEGER NOT NULL,

    PRIMARY KEY (question_id, subcategory_id),
    FOREIGN KEY (question_id) REFERENCES question(id) ON DELETE CASCADE,
    FOREIGN KEY (subcategory_id) REFERENCES subcategory(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE question_subcategory;
-- +goose StatementEnd

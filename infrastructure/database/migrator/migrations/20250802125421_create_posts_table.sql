-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE escalateservice.posts (
    post_id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    score INT DEFAULT 0,
    CONSTRAINT fk_username FOREIGN KEY (username)
        REFERENCES escalateservice.users (username)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE escalateservice.posts;
-- +goose StatementEnd

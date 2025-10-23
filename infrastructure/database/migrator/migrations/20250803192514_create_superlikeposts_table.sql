-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE escalateservice.superlikePosts (
    username VARCHAR(255) NOT NULL,
    post_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (username, post_id),
    CONSTRAINT fk_superlikepost_username FOREIGN KEY (username)
        REFERENCES escalateservice.users (username)
        ON DELETE CASCADE,
    CONSTRAINT fk_superlikepost_post_id FOREIGN KEY (post_id)
        REFERENCES escalateservice.posts (post_id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE escalateservice.superlikePosts;
-- +goose StatementEnd

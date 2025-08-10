-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE escalateservice.follows (
    follower VARCHAR(255) NOT NULL,
    followee VARCHAR(255) NOT NULL,
    PRIMARY KEY (follower, followee),
    CONSTRAINT fk_follow_follower FOREIGN KEY (follower)
        REFERENCES escalateservice.users (username)
        ON DELETE CASCADE,
    CONSTRAINT fk_follow_followee FOREIGN KEY (followee)
        REFERENCES escalateservice.users (username)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE escalateservice.follows;
-- +goose StatementEnd

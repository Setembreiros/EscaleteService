-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE escalateservice.reviews (
    review_id VARCHAR(255) PRIMARY KEY,
    post_id VARCHAR(255) NOT NULL,
    reviewer VARCHAR(255) NOT NULL,
    rating INT NOT NULL, -- Pode ser de 0 a 5

    CONSTRAINT fk_post_id FOREIGN KEY (post_id)
        REFERENCES escalateservice.posts (post_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_reviewer FOREIGN KEY (reviewer)
        REFERENCES escalateservice.users (username)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE escalateservice.reviews;
-- +goose StatementEnd

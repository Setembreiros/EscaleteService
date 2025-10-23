-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE escalateservice.users (
    username VARCHAR(255) PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE escalateservice.users;
-- +goose StatementEnd

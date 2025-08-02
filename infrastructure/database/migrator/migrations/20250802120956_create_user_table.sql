-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE escalateservice.users (
    username VARCHAR(255) NOT NULL  
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE escalateservice.users;
-- +goose StatementEnd

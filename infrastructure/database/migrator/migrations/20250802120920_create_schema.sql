-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SCHEMA escalateservice;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP SCHEMA escalateservice;
-- +goose StatementEnd

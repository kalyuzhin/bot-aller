-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    username TEXT PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd

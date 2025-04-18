-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats
(
    id SERIAL PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd

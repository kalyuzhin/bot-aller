-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats
(
    id SERIAL PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chats;
-- +goose StatementEnd

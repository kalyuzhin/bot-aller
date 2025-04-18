-- +goose Up
-- +goose StatementBegin
CREATE TABLE users_chats
(
    chat_id int,
    user_id int,
    FOREIGN KEY (chat_id) REFERENCES chats(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd

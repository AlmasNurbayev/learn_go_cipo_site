-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, email, password) VALUES ('admin', 'admin', 'admin');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users where name = 'admin';
-- +goose StatementEnd

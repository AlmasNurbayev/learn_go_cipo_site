-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
 id SERIAL PRIMARY KEY,
 email VARCHAR NOT NULL,
 name VARCHAR,
 password VARCHAR,
 salt VARCHAR,
 role VARCHAR,
 changed_at TIMESTAMPZ NOT NULL, 
 created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
 id BIGSERIAL PRIMARY KEY,
 email VARCHAR NOT NULL,
 name VARCHAR,
 password VARCHAR,
 salt VARCHAR,
 role VARCHAR,
 changed_at TIMESTAMPTZ, 
 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

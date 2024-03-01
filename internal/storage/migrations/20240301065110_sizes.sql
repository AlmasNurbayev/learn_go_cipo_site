-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sizes(
 id BIGSERIAL PRIMARY KEY,
 id_1c VARCHAR NOT NULL,
 name_1c VARCHAR NOT NULL,
 registrator_id BIGINT NOT NULL REFERENCES registrators (id),
 changed_at TIMESTAMPTZ, 
 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sizes;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS countries (
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
DROP TABLE countries;
-- +goose StatementEnd

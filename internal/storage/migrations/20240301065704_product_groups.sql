-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product_groups (
 id BIGSERIAL PRIMARY KEY,
 id_1c VARCHAR NOT NULL,
 name_1c VARCHAR NOT NULL,
 registrator_id BIGINT NOT NULL REFERENCES registrators (id),
 changed_at TIMESTAMPTZ, 
 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

 CONSTRAINT product_groups_id_1c UNIQUE (id_1c)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product_groups;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stores(
 id SERIAL PRIMARY KEY,
 id_1c VARCHAR NOT NULL,
 name_1c VARCHAR NOT NULL,
 address VARCHAR,
 registrator_id INT NOT NULL,
 link_2gis VARCHAR,
 phone VARCHAR,
 city VARCHAR,
 image_path VARCHAR,
 public BOOLEAN NOT NULL DEFAULT FALSE,
 changed_at TIMESTAMPZ NOT NULL, 
 created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table stores;
-- +goose StatementEnd

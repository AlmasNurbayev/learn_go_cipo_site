-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS registrators
(
 id BIGSERIAL PRIMARY KEY,
 operation_date TIMESTAMPTZ NOT NULL,
 name_folder VARCHAR NOT NULL,
 name_file VARCHAR NOT NULL,
 user_id BIGINT NOT NULL REFERENCES users (id),
 date_schema TIMESTAMPTZ NOT NULL,
 id_catalog VARCHAR NOT NULL,
 id_class VARCHAR NOT NULL,
 name_catalog VARCHAR NOT NULL,
 name_class VARCHAR NOT NULL,
 ver_schema VARCHAR NOT NULL,
 is_only_change BOOLEAN NOT NULL,
 changed_at TIMESTAMPTZ, 
 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table registrators;
-- +goose StatementEnd
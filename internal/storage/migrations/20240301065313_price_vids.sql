-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS price_vids (
 id BIGSERIAL PRIMARY KEY,
 id_1c VARCHAR NOT NULL,
 name_1c VARCHAR NOT NULL,
 registrator_id BIGINT NOT NULL REFERENCES registrators (id),
 is_active BOOLEAN NOT NULL DEFAULT FALSE,
 active_change_date TIMESTAMPTZ NOT NULL,
 changed_at TIMESTAMPTZ, 
 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

 CONSTRAINT price_vids_id_1c UNIQUE (id_1c)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE price_vids;
-- +goose StatementEnd

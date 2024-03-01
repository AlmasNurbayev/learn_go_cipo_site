-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS image_registry (
  id BIGSERIAL PRIMARY KEY,

  main BOOLEAN NOT NULL DEFAULT FALSE,
  main_change_at TIMESTAMPTZ NOT NULL, 
  resolution VARCHAR,
  size INT NOT NULL,
  full_name VARCHAR NOT NULL,
  name VARCHAR NOT NULL,

  path VARCHAR NOT NULL,
  operation_date TIMESTAMPTZ NOT NULL,
  active BOOLEAN NOT NULL DEFAULT FALSE,
  active_change_at TIMESTAMPTZ NOT NULL,
  product_id BIGINT NOT NULL REFERENCES products (id) NOT NULL,
  registrator_id BIGINT NOT NULL REFERENCES registrators (id),

  changed_at TIMESTAMPTZ, 
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS search_image
ON image_registry (product_id, active, main )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE image_registry;
DROP INDEX search_image;
-- +goose StatementEnd

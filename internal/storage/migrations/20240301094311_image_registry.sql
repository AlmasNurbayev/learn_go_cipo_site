-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS image_registry (
  id BIGSERIAL PRIMARY KEY,

  is_main BOOLEAN NOT NULL DEFAULT FALSE,
  resolution VARCHAR,
  size INT NOT NULL,
  full_name VARCHAR NOT NULL,
  name VARCHAR NOT NULL,

  path VARCHAR NOT NULL,
  operation_date TIMESTAMPTZ NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT FALSE,
  
  product_id BIGINT NOT NULL REFERENCES products (id) NOT NULL,
  registrator_id BIGINT NOT NULL REFERENCES registrators (id),

  changed_at TIMESTAMPTZ, 
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS search_image
ON image_registry (product_id, is_active, is_main )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE image_registry;
DROP INDEX search_image;
-- +goose StatementEnd

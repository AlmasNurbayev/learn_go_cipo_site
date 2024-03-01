-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS qnt_price_registry (
  id BIGSERIAL PRIMARY KEY,

  sum DECIMAL NOT NULL,
  qnt DECIMAL NOT NULL,
  operation_date TIMESTAMPTZ NOT NULL,
  discount_percent DECIMAL,
  discount_begin TIMESTAMPTZ,
  discount_end TIMESTAMPTZ,

  store_id BIGINT NOT NULL REFERENCES stores (id) NOT NULL,
  product_id BIGINT NOT NULL REFERENCES products (id) NOT NULL,
  price_vid_id BIGINT NOT NULL REFERENCES price_vids (id) NOT NULL,  
  size_id BIGINT NOT NULL REFERENCES sizes (id),
  registrator_id BIGINT NOT NULL REFERENCES registrators (id) NOT NULL,
  product_group_id BIGINT NOT NULL REFERENCES product_groups (id) NOT NULL,
  vid_modeli_id BIGINT NOT NULL REFERENCES vids (id),

  size_name_1c VARCHAR,
  product_name VARCHAR NOT NULL,
  product_created_at TIMESTAMPTZ,

  changed_at TIMESTAMPTZ, 
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS search_qnt_price
ON qnt_price_registry (registrator_id, product_id, product_created_at )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE qnt_price_registry;
DROP INDEX search_qnt_price;
-- +goose StatementEnd

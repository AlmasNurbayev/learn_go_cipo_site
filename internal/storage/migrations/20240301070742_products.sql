-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
  id BIGSERIAL PRIMARY KEY,
  id_1c VARCHAR NOT NULL,
  name VARCHAR NOT NULL,
  name_1c VARCHAR,

  product_group_id BIGINT NOT NULL REFERENCES product_groups (id),
  product_vid_id BIGINT NOT NULL REFERENCES product_vids (id),
  registrator_id BIGINT NOT NULL REFERENCES registrators (id),
  brand_id BIGINT REFERENCES brands (id),
  country_id BIGINT REFERENCES countries (id),
  vid_id BIGINT REFERENCES vids (id),

  artikul TEXT NOT NULL,
  base_ed VARCHAR NOT NULL,
  description VARCHAR,
  material_inside VARCHAR,
  material_podoshva VARCHAR,
  material_up VARCHAR,
  sex INT,
  product_folder VARCHAR,
  main_color VARCHAR,
  public_web BOOLEAN,
  changed_at TIMESTAMPTZ, 
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS id_reg
ON products (id, registrator_id, id_1c)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
DROP INDEX id_reg;
-- +goose StatementEnd

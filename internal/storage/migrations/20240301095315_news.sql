-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS news (
  id BIGSERIAL PRIMARY KEY,

  title VARCHAR NOT NULL,
  data TEXT NOT NULL,
  image_path VARCHAR,

  changed_at TIMESTAMPTZ, 
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE news;
-- +goose StatementEnd

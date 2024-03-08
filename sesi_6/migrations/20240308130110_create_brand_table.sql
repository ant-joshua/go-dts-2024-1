-- +goose Up
-- +goose StatementBegin
CREATE TABLE brands (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS brands;
-- +goose StatementEnd

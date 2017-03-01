
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE companies (
  id bigserial NOT NULL,
  name varchar(255) NOT NULL,
  phone varchar(50) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE companies;

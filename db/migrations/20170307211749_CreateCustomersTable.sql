
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE customers (
  id bigserial NOT NULL,
  company_id bigint NOT NULL,
  name varchar(255) NOT NULL,
  phone varchar(50) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE customers;

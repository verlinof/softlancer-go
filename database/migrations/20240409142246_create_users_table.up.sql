CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  born_date DATE
)

-- Command to run migrations
-- migrate -path database/migrations -database "mysql://root:@tcp(localhost:3306)/go_restful" up

-- How to add Migrations
-- migrate create -ext sql -dir database/migrations create_name_table
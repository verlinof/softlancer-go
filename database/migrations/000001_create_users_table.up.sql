CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL
)

-- migrate -path database/migrations -database "mysql://root:@tcp(localhost:3306)/go_restful" up
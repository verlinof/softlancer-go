drop TABLE if exists users

-- Command to revert migrations
-- migrate -path database/migrations -database "mysql://root:@tcp(localhost:3306)/golang_restful" down
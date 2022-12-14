#window
migrate -source file:db\migration -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable up 1

#docker down 2
migrate -source file:db\migration -database postgres://postgres:admin@localhost:5433/postgres?sslmode=disable down 2

#mac
migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/postgres?sslmode=disable" -verbose up 1
g

#test
go test -v .\internal\repository\user\ -run TestGetUsers
#test mac 
go test -v ./internal/repository/user/ -run TestGetUsers
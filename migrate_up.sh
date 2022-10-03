#window
migrate -source file:db\migration -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable up 1
migrate -source file:db\migration -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable down 1

#mac
migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/postgres?sslmode=disable" -verbose up 1


#test
go test -v .\internal\repository\user\ -run TestGetUsers
#test mac 
go test -v ./internal/repository/user/ -run TestGetUsers
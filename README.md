# An-Order-Management
Golang Practice REST API

## Order-MG
An Order Management

## Description
This is my REST API Golang practice with a basic CRUD list of features on both User and Order. 

## Technologies
- Golang 1.19
- Gochi 
- GORM 
- SonyFlake
- Logrus
- Mockery
- Testify
- PostgreSQL
- Docker

## Executing Program
### to migrate the db:  
migrate -source file:db\migration -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable up 1

### to run the program: 
go run .\cmd\served\main.go

### by docker: 
docker-compose up db app migrate_up

## Postman test
https://www.getpostman.com/collections/a09c1de3c409178f1f44

####

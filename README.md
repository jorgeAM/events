# EVENTS 🚀
First at all run `docker-compose up -d`

## Event Service
To run microservice just do the following:
1. go to the root of event service
2. run `go mod download`
3. `go run main.go`


## Booking Service
To run microservice just do the following:
1. go to the root of booking service
2. run `go mod download`
3. run migrations, at the moment migrations only works with postgres

to run migrations just write these command `goose -dir "migrations" postgres "postgres_url" up -v`

4. `go run main.go`

PS: your feedbacks are welcome 

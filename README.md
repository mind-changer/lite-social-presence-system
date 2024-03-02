# lite-social-presence-system

## Docker
docker compose build
docker compose run
go to localhost:5050
login with following credentials
username- admin@admin.com
password- admin

Create a new server
with name as postgres
and host, username, password as postgres

Run the loaddb.go script inside the web container using docker exec
go run ./scripts/loaddb.go

Open postman
Now you can run any http calls on localhost:80 and grpc calls on localhost:81

## Kubernetes
# lite-social-presence-system

## Docker

Steps:

start docker
docker compose build
docker compose run

(go to localhost:5050
login with following credentials
username- admin@admin.com
password- admin
Create a new server
with name as postgres
and host, username, password as postgres)

OR 

(docker compose exec postgres sh
psql -U postgres -d mydb)

docker compose exec web sh
Run the loaddb.go script inside the web container using docker exec
go run ./scripts/loaddb.go

Open postman
Now you can run any http calls on localhost:80 and grpc calls on localhost:81

## Kubernetes

Steps:

start docker
minikube start
kubectl apply -f deployment.yaml
kubectl get all

kubectl exec -it postgres-859fb969d4-9tkqk bash
psql -U postgres -d mydb

kubectl exec -it pod/webapp-5cc44dfd5f-z5m2v -- /bin/bash
go run ./scripts/loaddb.go

kubectl port-forward deployment.apps/webapp 80:80

kubectl port-forward deployment.apps/webapp 81:81

Now you can run any http calls on localhost:80 and grpc calls on localhost:81
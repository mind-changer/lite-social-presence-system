# lite-social-presence-system

Clone the source code using:

1. git clone https://github.com/mind-changer/lite-social-presence-system.git
2. cd lite-social-presence-system

## Docker

Steps:

1. Start docker daemon
2. Open cmd
3. docker compose build
4. docker compose up
5. Enter the shell of postgres container: docker compose exec postgres sh
6. psql -U postgres -d mydb

OR 

you can go to http://localhost:5050
login with following credentials
username- admin@admin.com
password- admin
Create a new server with name as postgres and fill connection details host, username, password as postgres

7. Enter the shell of web container: docker compose exec web sh
8. Run the scripts/loaddb.go script inside the web container to load initial dummy db records: go run ./scripts/loaddb.go

Open postman or you can use curl and run any http calls on localhost:80 and grpc calls on localhost:81

## Kubernetes

Steps:

1. Start docker daemon
2. Start minikube: minikube start
3. Apply the kubernetes config: kubectl apply -f deployment.yaml
4. Grab the postgres pod name using: kubectl get all
5. Enter the bash of the postgres pod. Run this command by replacing "postgres-859fb969d4-9tkqk" with the pod name: kubectl exec -it postgres-859fb969d4-9tkqk -- /bin/bash
6. psql -U postgres -d mydb
7. Enter the bash of the webapp pod. Run this command by replacing "pod/webapp-5cc44dfd5f-z5m2v" with the pod name: kubectl exec -it pod/webapp-5cc44dfd5f-z5m2v -- /bin/bash
8. Run the scripts/loaddb.go script inside the webapp pod to load initial dummy db records: go run ./scripts/loaddb.go
9. Forward the http and grpc ports in two separate terminals 
kubectl port-forward deployment.apps/webapp 80:80
kubectl port-forward deployment.apps/webapp 81:81

Now you can run any http calls on localhost:80 and grpc calls on localhost:81 using postman or curl

## Postman

You can import the http_lite_social_presence_system.postman_collection.json file into your postman to test out the HTTP APIs
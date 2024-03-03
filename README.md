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

## API testing

Friends

1. view friends
![Screenshot (14)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/1388b6b1-4d8b-4950-ad79-441c7081f672)
2. send friend request
![Screenshot (15)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/0aa6c4a6-73b2-4f5e-9cb5-e31aba944da3)
![Screenshot (16)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/84cc62f5-0694-4ba1-99d2-07b6d428f4e1)

4. accept friend request
![Screenshot (17)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/6f9c533a-f744-4d96-9955-f3525378523c)
![Screenshot (18)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/6f2404f1-af42-4f8c-9c6f-068785e07b67)

6. remove friend
![Screenshot (19)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/fbbb4d92-a3a1-48d2-be89-9064099f7e34)
![Screenshot (20)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/163d23fe-9645-4377-8076-1bc9c4d08d66)

8. send friend request again
![Screenshot (21)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/c80acb45-ce43-4e68-bdfa-e00a8ae7ec4b)

10. reject friend request
![Screenshot (22)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/a635610a-76ec-49f4-ba32-1612b5a04063)


Party

1. create party
2. invite friend to party
3. accept invitation
4. leave party
5. send invitation again
6. reject invitation
7. send invitation again
8. accept invitation 
9. kick member

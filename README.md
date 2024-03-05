# lite-social-presence-system

Clone the source code using:

1. `git clone https://github.com/mind-changer/lite-social-presence-system.git`
2. `cd lite-social-presence-system`

## Docker

Steps:

1. Start docker daemon
2. Open cmd
3. Build docker image: `docker compose build`
4. Run docker compose: `docker compose up`
5. Enter the shell of postgres container: `docker compose exec postgres sh`
6. Run the following command in the postgres container shell to connect to mydb: `psql -U postgres -d mydb` OR Alternatively, you can go to http://localhost:5050 and login with username as `admin@admin.com` and password as `admin` and then create a new server with name as `postgres` and fill connection details with host, username, password as `postgres`
7. Exit from the postgres container shell. Now, enter the shell of web container: `docker compose exec web sh`
8. Run the scripts/loaddb.go script inside the web container to load initial dummy db records: `go run ./scripts/loaddb.go`

Open postman or you can use curl and run any http calls on localhost:80 and grpc calls on localhost:81

## Kubernetes

Steps:

1. Start docker daemon
2. Start minikube: minikube start
3. Apply the kubernetes config file: `kubectl apply -f deployment.yaml`
4. Grab the postgres pod name using: `kubectl get all`
5. Enter the bash of the postgres pod. Run this command by replacing "postgres-859fb969d4-9tkqk" with the pod name: `kubectl exec -it postgres-859fb969d4-9tkqk -- /bin/bash`
6. Run the following command in the postgres pod bash to connect to mydb:`psql -U postgres -d mydb`
7. Exit from the postgres pod bash. Now, enter the bash of the webapp pod. Run this command by replacing "pod/webapp-5cc44dfd5f-z5m2v" with the pod name: `kubectl exec -it pod/webapp-5cc44dfd5f-z5m2v -- /bin/bash`
8. Run the scripts/loaddb.go script inside the webapp pod to load initial dummy db records: `go run ./scripts/loaddb.go`
9. Forward the http and grpc ports in two separate terminals 
`kubectl port-forward deployment.apps/webapp 80:80`
`kubectl port-forward deployment.apps/webapp 81:81`

Now you can run any http calls on localhost:80 and grpc calls on localhost:81 using postman or curl

## Postman

You can import the http_lite_social_presence_system.postman_collection.json file into your postman to test out the HTTP APIs

## APIs

### Update user online status API
```
curl --location --request PUT 'localhost:80/lite-social-presence-system/users/{user-id}/status' \
--header 'Content-Type: application/json' \
--data '{
    "userStatus":"offline"
}'
```

###  View Friends API
```
curl --location 'localhost:80/lite-social-presence-system/users/{user-id}/friends'
```

### Send friend request API
```
curl --location 'localhost:80/lite-social-presence-system/users/{user-id}/friends' \
--header 'Content-Type: application/json' \
--data '{
    "userId":"friendId"
}'
```

### Accept friend request API
```
curl --location --request PATCH 'localhost:80/lite-social-presence-system/users/{user-id}/friend-requests/{requester-id}'
```

### Reject friend request API
```
curl --location --request DELETE 'localhost:80/lite-social-presence-system/users/{user-id}/friend-requests/{requester-id}'
```

### Remove friend API
```
curl --location --request DELETE 'localhost:80/lite-social-presence-system/users/{user-id}/friends/{friend-id}'
```

### Create party API
```
curl --location --request POST 'localhost:80/lite-social-presence-system/users/{user-id}/parties'
```

### Send party invitation API
```
curl --location 'localhost:80/lite-social-presence-system/users/{user-id}/parties/{party-id}/member-invitations' \
--header 'Content-Type: application/json' \
--data '{
    "userId":"hillock123"
}'
```

### Accept party invitation API
```
curl --location --request PATCH 'localhost:80/lite-social-presence-system/users/{user-id}/party-invitations/{party-id}'
```

### Reject party invitation API
```
curl --location --request DELETE 'localhost:80/lite-social-presence-system/users/{user-id}/party-invitations/{party-id}'
```

### Leave party API
```
curl --location --request DELETE 'localhost:80/lite-social-presence-system/users/{user-id}/joined-parties/current/{party-id}'
```

### Kick party member API
```
curl --location --request DELETE 'localhost:80/lite-social-presence-system/users/{user-id}/parties/{party-id}/members/{member-id}'
```

## gRPC
### Get user online status
check the tested example below https://github.com/mind-changer/lite-social-presence-system/blob/main/README.md#grpc-services

### Get real time party members
check the tested example below https://github.com/mind-changer/lite-social-presence-system/blob/main/README.md#grpc-services

## Testing(screenshots might take time to load)
### HTTP RESTful APIs

#### Friends

You can view friends of userid "bnb"
![Screenshot (14)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/1388b6b1-4d8b-4950-ad79-441c7081f672)

User "bnb" sends a friend request to user "supergamer"
![Screenshot (15)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/0aa6c4a6-73b2-4f5e-9cb5-e31aba944da3)

We see a new record in the friend_requests table
![Screenshot (16)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/84cc62f5-0694-4ba1-99d2-07b6d428f4e1)

User "supergamer" now accepts bnb's friend request
![Screenshot (17)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/6f9c533a-f744-4d96-9955-f3525378523c)

We can see supergamer has been added as a friend of bnb
![Screenshot (18)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/6f2404f1-af42-4f8c-9c6f-068785e07b67)

Now bnb wants to remove supergamer from being a friend
![Screenshot (19)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/fbbb4d92-a3a1-48d2-be89-9064099f7e34)

We see the record of friends has been updated
![Screenshot (20)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/163d23fe-9645-4377-8076-1bc9c4d08d66)

Now bnb sends a friend request again
![Screenshot (21)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/c80acb45-ce43-4e68-bdfa-e00a8ae7ec4b)

supergamer rejects bnb's friend request
![Screenshot (22)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/a635610a-76ec-49f4-ba32-1612b5a04063)

Friends table looks like this now
![Screenshot (23)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/105cf5f2-059b-40dc-a963-e367886b80d5)

Friend request deleted
![Screenshot (24)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/2a8dfa0b-a7df-49a1-b836-77dc845610d4)

#### Party

User "bnb" creates a new party. The party ID is returned in the response.
![Screenshot (25)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/516699b1-e71e-440a-bc0c-c078cc331901)

Parties table shows a new party with it's ID and owner
![Screenshot (26)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/d06fcdd4-8447-4830-8f27-1cb852201682)

User "bnb" invites his friend "hillock123" to his party
![Screenshot (27)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/ac9e4481-a96a-4d2d-9e72-f7f09b86f17f)

Party invitations table looks like
![Screenshot (28)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/0683f9f9-67d7-4c33-8eab-721795942c26)

Party members table looks like
![Screenshot (29)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/9b2a2157-7925-4e91-999b-49bd2c2c8467)

User "hillock123" accepts the party invitation and joins bnb's party
![Screenshot (30)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/03cb9c0d-37c6-4065-b375-5987c7d3fe48)

Party members table looks like
![Screenshot (31)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/cd82c2ec-d22d-48ad-ac00-c89f635b79f0)

Party invitations table looks like
![Screenshot (32)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/018e936a-134d-4afb-b2ce-652ccf66680f)

User "hillock123" leaves the party
![Screenshot (33)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/d6c890b6-ef76-4d0d-beec-5b7767dc8746)

Party members table looks like
![Screenshot (34)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/50f50b91-9146-420e-a246-8f7fd1057545)

User "bnb" sends a party invitation again to "hillock123"
![Screenshot (35)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/9b688c2c-47e3-4bdb-a2b6-10f1b7bd82e6)

Party invitations table looks like
![Screenshot (36)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/4fcad6a4-23e8-4093-8ee6-f93a4132fa61)

User "hillock123" rejects the party invitation
![Screenshot (37)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/b422b906-6aed-4c91-8fa2-0497b62023b1)

Party invitations table looks like
![Screenshot (38)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/898db524-607b-46c7-a134-48490bf6762c)

Party members table looks like
![Screenshot (39)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/b86e3bf2-892c-4d4c-984b-dfe1654b7b02)

User "bnb" sends a party invitation again to "hillock123"
![Screenshot (40)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/d915a0b9-4ed3-4aa2-8465-30d9253c190b)

User "hillock123" accepts the party invitation and joins bnb's party
![Screenshot (41)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/2938b847-93b1-4900-86a4-8c83925a9c7f)

User "bnb" kicks out "hillock123" out of his party
![Screenshot (42)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/9a0d5f9d-5562-48d1-bb05-99663c029b4b)

Party members table looks like
![Screenshot (43)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/a5c273ef-fe95-48d1-8f53-cd6771f21cda)

Party invitation table looks like
![Screenshot (44)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/d0690e03-2eff-44db-a40e-81131ca5370a)

### gRPC services

#### Get user online status

Get User "hillock123" online status. The value of status is checked and sent every 5 seconds to the client. Currently he's online.
![Screenshot (45)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/d0b5449b-b885-4235-b5da-c5a70557a8ad)

![Screenshot (46)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/3e626785-d3c0-4c45-b2a5-c5d5dd5f0580)

Now update user online status to "offline"
![Screenshot (47)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/40a8ddfa-62f4-4dda-828b-c20beef42144)

User's real time online status changes from "online" to "offline"
![Screenshot (48)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/8f86f89f-5141-4bcc-b446-8ca64706724e)

#### Get real time party member status

Get the party members of a party real time, using party ID. The latest party members are checked and sent every 5 seconds to the client.
![Screenshot (49)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/3683d777-8ae6-46ce-b77d-415ff1e15241)

Invite a member to the party
![Screenshot (50)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/e1406ae3-f360-4d84-b910-0d441b738bd3)

User accepts party invitation and joins the party
![Screenshot (51)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/72e1e289-d3e9-4e5c-9157-29a931c617d5)

The party member list is updated
![Screenshot (52)](https://github.com/mind-changer/lite-social-presence-system/assets/43662445/ee31a496-5c45-42db-8d33-f62f389ad510)


## Task 6

## Description

Migrating from the REST approach to gRPC + Gateway.

1. Take code with the latest changes from [Task 5](../task_5).
2. Create a new repo with the same name as a previous and add `-grpc` postfix(Example: `golang-training-theater-grpc`).  
3. Create proto files for your models and services with required methods from your previous implementation (CRUD).
4. Generate grpc and gateway from the proto files.
5. Rewrite your code to gateway and grpc server layers.
7. Add go modules with vendoring.
8. Add logger using logrus golang library.
9. Add dockerfiles (gateway, grpc server, db).
10. docker-compose that allows you run application locally.
11. Add readme file with the following paragraphs.
    - project description
    - requirements for running application locally
    - table with paths
    - how to generate go code from protobuf
    - how to run application via docker-compose

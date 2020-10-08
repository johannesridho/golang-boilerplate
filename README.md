# Golang Boilerplate

## How to Run
1. Run the command `make run` to run the server on localhost:8080.
2. Use a grpc client to send a request to the server. For example, using grpc_cli, 
you can use these command:
    - `grpc_cli ls localhost:8080 -l` to list all the services available
    - `grpc_cli call 127.0.0.1:8080 GetUsers ""` to call GetUsers rpc

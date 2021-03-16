# gRPC + Authentication + Gateway

This repo aims to provide an example where a Go Server accepts both REST and RPC requests.
The RPC communication between Server and Client is made with [gRPC](https://grpc.io/) while the other has been made possible 
thanks [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/). 

The server exposes two methods defined in the `proto/todo/todo.proto`. The first one will add a Todo to the server and the second one
will get all the added Todos.

An authentication string has to be provided in order to interact with the server. The gRPC client uses `grpc.WithPerRPCCredentials` 
while the REST client have to provide the key in the header request (`todo-grpc-gateway`).

### Prerequisites

Install the dependencies in order to run the program locally. 
The following command will install also `protoc` to generate Go stubs for your own further development.

```bash
$ go mod download
```

### Running

The `main.go` will run the gRPC server (listening on port: `50051`) and the gRPC-Gateway (listening on port `8080`).

```bash
$ go run main.go
```

### Testing

Send RPC request (authentication credentials are already set in the source code).

```bash
$ go run client/main.go
```

Send Add Todo REST request (authentication in the header).

```bash
$ curl -X POST \
  http://localhost:8080/api/todo \
  -H 'content-type: application/json' \
  -H 'todo-grpc-gateway: passcode' \
  -d '{"name": "todo1"}'
```

Send Get Todos REST request (authentication in the header).

```bash
$ curl -X GET \
  http://localhost:8080/api/todos \
  -H 'content-type: application/json' \
  -H 'todo-grpc-gateway: passcode'
```

### Customize

The next step is to define the methods you want to expose in `proto/todo/todo.proto`.
Once that is done, regenerate the files using the following command. 
This will mean you'll need to implement any functions in `server/server.go`, or else the build will fail since 
your struct won't be implementing the methods defined. 

```bash
$ protoc -I ./proto \
   --go_out ./proto --go_opt paths=source_relative \
   --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
   ./proto/todo/todo.proto
```

The authentication can be improved quite a lot though
# api-handler example
This repository shows an example of how you can create an API handler for you application. The API handler created HTTP routes, uses gin web framework to handler HTTP reqeusts in HTTP endpoints and then creates gRPC clients using google's grpc library to send a gRPC request to the microservice that has the gRPC server.

## Development

### Installation
1. Install go (go version go1.19.4 darwin/amd64)
  - Set your paths accordingly
2. Install protobuf and grpc (use @latest if needed)
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@vv1.30.0
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```

### Dependencies
These exists in go.mod file, have a look there and `go get <library>` accordingly when you start a new repo.

### Workflow
- When implementing API or gateways or services that uses grpc
- Follow the folder structure to for the `.proto` files
  - ./pkg/<service-name>/pb/*.proto
- Use the same Makefile command to generate `pb.go` files
  - This will generate and output the files in the same directory as .proto files

## Web Framework choice
We chose to use Gin, initially wanted to use go-chi but realised that it doens't have a context like gin does to handle HTTP requests. 

### Common third party libraries that people tend to use
1. gin - HTTP request handler
2. go-chi - HTTP request handler
3. viper - configuration management/ solution
  - See more: https://github.com/spf13/viper
4. grpc - protobuf communications

## Standard Go application architecture
Inspiration: https://levelup.gitconnected.com/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30
1. Microservices that communicates via gRPC
2. API Gateway - receives HTTP, communicate with  microservices via gRPC
  - This would be what the frontend communicates with

## Poject structure
menumeister-api/
├── cmd/
│   ├── main.go
├── pkg/
│   ├── config/
│   │   ├── envs/
│   │   │   └── dev.env
│   │   └── config.go
│   │── ping/
│   │   ├── pb/
│   │   │   └── ping.proto
│   │   ├── routes/
│   │   │   └── ping.go
│   │   ├── routes.go
│   │   └── client.go
├── go.mod
├── go.sum
├── Makefile
└── Other docker stuff, add as needed

This is the normal conventional directory structure, you could include an internal directory in the same level as cmd, this directory will house logic that is internal to this application only.

### Key notes
1. cmd - houses all files that executes the application. can include seed and migration as well, any sort of execution on the applciation itself
2. pkg - houses any business logic that may be required by ur other external microservices
  - May want to shift config out
3. ping directory in pkg is an example microservice that this api gateway wants to call
  - pb directory houses the protobuf message strucutres that is used in gRPC communication


# Helps automate tasks

# generate go code from all .proto files found in the specified directory
# since we structured the project to house all proto files in pkg/**/pb, this 
# will help us automate go file generation by doing `make proto`
# --go_out flag specifies that generated files should be placed in the same dir as proto files
proto:
	protoc --proto_path=./pkg --go-grpc_out=. --go_out=paths=import:. ./pkg/**/pb/*.proto

server:
	go run cmd/main.go

docker:
	docker build --tag menumeister-api .
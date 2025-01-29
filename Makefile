setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/wire/cmd/wire@latest

protoc:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

gqlgen:
	go run github.com/99designs/gqlgen generate

wire:
	wire

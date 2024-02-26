build:
	go build -o bin/main.exe cmd/main.go

run: build
	go run bin/main.exe

proto:
	protoc --go_out=./pkg --proto_path=./submodules/grpc-bookshelf-proto/proto ./submodules/grpc-bookshelf-proto/proto/book.proto
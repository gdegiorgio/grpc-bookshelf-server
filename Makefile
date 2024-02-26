build:
	go build -o bin/main.exe cmd/main.go

run: build
	go run bin/main.exe
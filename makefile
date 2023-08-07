proto:
	protoc --go_out=. --go-grpc_out=. ./proto/*.proto

wire:
	cd di && wire

run:
	go run cmd/main.go


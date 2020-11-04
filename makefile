.PHONY: gen clean server client

gen:
	protoc --proto_path=proto/ proto/*.proto --go_out=plugins=grpc:pb
clean:
	rm pb/*.go
server:
	go run grpc/server/main.go
client:
	go run grpc/client/main.go